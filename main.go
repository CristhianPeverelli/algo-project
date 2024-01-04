package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type mattoncino struct {
	Alpha string // bordo α
	Beta  string // bordo β
	Sigma string // nome
}

type fila struct {
	Mattoncini        []*mattoncino
	BordiSinistri     map[string]string
	BordiDestri       map[string]string
	SimboliMattoncini map[string]byte
}

type scatola struct {
	Mattoncini map[string]*mattoncino
}

type gioco struct {
	Scatola      scatola
	FileDisposte *[]fila
}

func main() {
	filePath := "C:\\Users\\cpeverelli\\Documents\\GitHub\\algo-project\\input1.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scatola := scatola{
		Mattoncini: make(map[string]*mattoncino),
	}
	gioco := gioco{
		Scatola:      scatola,
		FileDisposte: &[]fila{},
	}
	//scanner := bufio.NewScanner(os.Stdin)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("Start: ") //TODO: da rimuovere
	for _, linea := range input {
		elementi := strings.Fields(linea)
		switch elementi[0] {
		case "m":
			alpha := elementi[1]
			beta := elementi[2]
			sigma := elementi[3]
			inserisciMattoncino(gioco, alpha, beta, sigma)
			break
		case "s":
			stampaMattoncino(gioco, linea[2:])
			break
		case "d":
			disponiFila(gioco, linea[2:])
			break
		case "S":
			stampaFila(gioco, linea[2:])
			break
		case "e":
			eliminaFila(gioco, linea[2:])
			break
		case "f":
			alpha := elementi[1]
			beta := elementi[2]
			disponiFilaMinima(gioco.Scatola.Mattoncini, alpha, beta)
			break
		case "M":
			sigma := elementi[1]
			tau := elementi[2]
			sottostringaMassima(sigma, tau)
			break
		case "i":
			fmt.Println("Cerco:", linea[2:])
			indiceCacofonia(gioco, linea[2:])
			break
		case "c":
			break
		case "q":
			fmt.Println(":End")
			return
		}

	}

}

/*
inserisciMattoncino (α, β, σ)
Se esiste già un mattoncino di nome σ oppure se α `e uguale a β, non compie alcuna operazione.
Altrimenti, inserisce nella scatola il mattoncino definito dalla tripla (α, β, σ).
*/
func inserisciMattoncino(g gioco, alpha, beta, sigma string) {
	if alpha != beta && g.Scatola.Mattoncini[sigma] == nil {
		mattoncino := &mattoncino{Alpha: alpha, Beta: beta, Sigma: sigma}
		g.Scatola.Mattoncini[mattoncino.Sigma] = mattoncino
	}
}

/*
stampaMattoncino (σ)
Se non esiste alcun mattoncino di nome σ non compie alcuna operazione. Altrimenti, stampa il
mattoncino con nome σ, secondo il formato specificato nell’apposita sezione.
*/
func stampaMattoncino(g gioco, sigma string) {
	if g.Scatola.Mattoncini[sigma] != nil {
		m := g.Scatola.Mattoncini[sigma]
		fmt.Printf("%s: %s, %s\n", m.Sigma, m.Alpha, m.Beta)
	}
}

/*
disponiFila (±σ1, ±σ2, . . . , ±σn)
dove ± indica uno dei due simboli + o −. Verifica se nella scatola ci sono i mattoncini di nome
σ1, σ2, . . . , σn e se la sequenza di mattoncini ±σ1, ±σ2, . . . , ±σn costituisce una fila; in questo caso,
toglie dalla scatola i mattoncini che la comp
*/
func inserisciMattoncinoInFila(f *fila, m *mattoncino, standard bool) {
	f.Mattoncini = append(f.Mattoncini, m)
	if standard { //Simbolo +
		f.BordiSinistri[m.Sigma] = m.Alpha
		f.BordiDestri[m.Sigma] = m.Beta
		f.SimboliMattoncini[m.Sigma] = '+'
	} else { //Simbolo -
		f.BordiSinistri[m.Sigma] = m.Beta
		f.BordiDestri[m.Sigma] = m.Alpha
		f.SimboliMattoncini[m.Sigma] = '-'
	}
}

func compatibili(m1, m2 *mattoncino, s1, s2 byte) bool {
	if s1 == '+' {
		if s2 == '+' {
			if m1.Beta != m2.Alpha {
				return false
			}
		} else {
			if m1.Beta != m2.Beta {
				return false
			}
		}
	} else {
		if s2 == '+' {
			if m1.Alpha != m2.Alpha {
				return false
			}
		} else {
			if m1.Alpha != m2.Beta {
				return false
			}
		}
	}
	return true
}

func disponiFila(g gioco, listaNomi string) {

	nomi := strings.Fields(listaNomi)
	if len(nomi) == 0 {
		return
	}

	fila := &fila{Mattoncini: make([]*mattoncino, 0), BordiSinistri: make(map[string]string), BordiDestri: make(map[string]string), SimboliMattoncini: make(map[string]byte)}
	if len(nomi) == 1 {
		mattoncino := g.Scatola.Mattoncini[nomi[0][1:]]
		inserisciMattoncinoInFila(fila, mattoncino, nomi[0][0] == '+')
		*g.FileDisposte = append(*g.FileDisposte, *fila)
		return
	}

	mattonciniPresi := make(map[string]bool)
	for i := 0; i < len(nomi)-1; i++ {
		var mattoncino1 *mattoncino
		var mattoncino2 *mattoncino
		if !mattonciniPresi[nomi[i][1:]] {
			mattoncino1 = g.Scatola.Mattoncini[nomi[i][1:]]
			mattonciniPresi[mattoncino1.Sigma] = true
		}
		if mattoncino1 == nil {
			return
		}
		mattonciniPresi[mattoncino1.Sigma] = true
		if !mattonciniPresi[nomi[i+1][1:]] {
			mattoncino2 = g.Scatola.Mattoncini[nomi[i+1][1:]]
		}
		if mattoncino2 == nil {
			return
		}

		segno1 := nomi[i][0]
		segno2 := nomi[i+1][0]
		if !compatibili(mattoncino1, mattoncino2, segno1, segno2) {
			return
		}
		inserisciMattoncinoInFila(fila, mattoncino1, (nomi[i][0] == '+'))
	}
	var m *mattoncino
	if !mattonciniPresi[nomi[len(nomi)-1][1:]] {
		m = g.Scatola.Mattoncini[nomi[len(nomi)-1][1:]]
		mattonciniPresi[m.Sigma] = true
	}
	inserisciMattoncinoInFila(fila, m, nomi[len(nomi)-1][0] == '+')
	for mat := range mattonciniPresi {
		delete(g.Scatola.Mattoncini, mat)
	}
	*g.FileDisposte = append(*g.FileDisposte, *fila)
}

/*
stampaFila (σ)
Se non esiste alcun mattoncino di nome σ, oppure se il mattoncino di nome σ non appartiene
ad alcuna fila sul tavolo da gioco, non compie alcuna operazione. Altrimenti, stampa la fila cui
appartiene il mattoncino con nome σ, secondo il formato specificato nell’apposita sezione.
*/

func contiene(f fila, sigma string) bool {
	return f.BordiDestri[sigma] != ""
}

func stampaFila(g gioco, sigma string) {
	for _, f := range *g.FileDisposte {
		if contiene(f, sigma) {
			fmt.Println("(")
			for _, m := range f.Mattoncini {
				if f.SimboliMattoncini[m.Sigma] == '+' {
					fmt.Printf("%s: %s, %s\n", m.Sigma, m.Alpha, m.Beta)
				} else {
					fmt.Printf("%s: %s, %s\n", m.Sigma, m.Beta, m.Alpha)
				}
			}
			fmt.Println(")")
			return
		}
	}
}

func (f *fila) getNome() (nome string) {
	for _, m := range f.Mattoncini {
		nome += m.Sigma
	}
	return nome
}

/*
eliminaFila (σ)
Se non esiste alcun mattoncino di nome σ, oppure se il mattoncino di nome σ non appartiene ad
alcuna fila sul tavolo da gioco, non compie alcuna operazione. Altrimenti, sia F la fila cui appartiene
il mattoncino di nome σ. La fila F `e rimossa dal tavolo e tutti i mattoncini che la compongono
sono rimessi nella scatola
*/
func inserisciInScatola(g gioco, f fila) {
	for _, m := range f.Mattoncini {
		g.Scatola.Mattoncini[m.Sigma] = m
	}
}

func eliminaFila(g gioco, sigma string) {
	for i, f := range *g.FileDisposte {
		if contiene(f, sigma) {
			inserisciInScatola(g, f)
			if i == 0 {
				if len(*g.FileDisposte) > 1 {
					*g.FileDisposte = (*g.FileDisposte)[1:]
				} else {
					*g.FileDisposte = nil
				}
			} else if i == len(*g.FileDisposte)-1 {
				*g.FileDisposte = (*g.FileDisposte)[:i]
			} else {
				*g.FileDisposte = append((*g.FileDisposte)[:i], (*g.FileDisposte)[i+1:]...)
			}
			break
		}
	}
}

/*
disponiFilaMinima (α, β)
Crea e posiziona sul tavolo da gioco una fila di lunghezza minima da α a β. Tutti i mattoncini
della fila devono essere presi dalla scatola. Se non `e possibile creare alcuna fila da α a β, stampa il
messaggio: non esiste fila da α a β
*/
func findShortestPath(mattoncini map[string]*mattoncino, current *mattoncino, beta string, visited map[string]bool, currentPath *fila, shortestPath *fila) {
	if current.Beta == beta {
		if len(shortestPath.Mattoncini) == 0 || len(currentPath.Mattoncini) < len(shortestPath.Mattoncini) {
			shortestPath.Mattoncini = append([]*mattoncino{}, currentPath.Mattoncini...)
		}
		return
	}
	visited[current.Sigma] = true
	checkMattoncini := make([]*mattoncino, 0)
	checkSimboli := make([]byte, 0)
	for _, m := range mattoncini {
		if m.Alpha == current.Beta {
			checkMattoncini = append(checkMattoncini, m)
			checkSimboli = append(checkSimboli, '+')
		} else if m.Beta == current.Beta {
			checkMattoncini = append(checkMattoncini, m)
			checkSimboli = append(checkSimboli, '-')
		}
	}
	for _, nextMattoncino := range checkMattoncini {
		if !visited[nextMattoncino.Sigma] && (current.Beta == nextMattoncino.Alpha || current.Beta == nextMattoncino.Beta) {
			simbolo := byte('+')
			if current.Beta == nextMattoncino.Beta {
				simbolo = byte('-')
			}
			nextPath := &fila{
				Mattoncini:        append([]*mattoncino{}, currentPath.Mattoncini...),
				SimboliMattoncini: make(map[string]byte),
			}
			nextPath.SimboliMattoncini[nextMattoncino.Sigma] = simbolo
			fmt.Println("CONTROLLO:", nextMattoncino)
			findShortestPath(mattoncini, nextMattoncino, beta, visited, nextPath, shortestPath)
		}
	}
	visited[current.Sigma] = false
}

func disponiFilaMinima(mattoncini map[string]*mattoncino, alpha, beta string) *fila {
	startMattoncini := make([]*mattoncino, 0)
	startSimboli := make([]byte, 0)
	for _, m := range mattoncini {
		if m.Alpha == alpha {
			startMattoncini = append(startMattoncini, m)
			startSimboli = append(startSimboli, '+')
		} else if m.Beta == alpha {
			startMattoncini = append(startMattoncini, m)
			startSimboli = append(startSimboli, '-')
		}
	}
	if len(startMattoncini) == 0 {
		return nil
	}

	shortestPath := &fila{
		Mattoncini:        make([]*mattoncino, 0),
		SimboliMattoncini: make(map[string]byte),
	}

	visited := make(map[string]bool)
	for i, m := range startMattoncini {
		currentPath := &fila{
			Mattoncini:        []*mattoncino{m},
			SimboliMattoncini: make(map[string]byte),
		}
		currentPath.SimboliMattoncini[m.Sigma] = startSimboli[i]
		fmt.Println("PARTO DA: ", m)
		findShortestPath(mattoncini, m, beta, visited, currentPath, shortestPath)
	}

	fmt.Println("FILA MINIMA:")
	for _, m := range shortestPath.Mattoncini {
		fmt.Println(m.Sigma)
	}

	return shortestPath
}

/*
sottostringaMassima (σ, τ )
Stampa su una nuova riga una sottostringa massima ρ di σ e τ (se ρ `e la stringa nulla, stampa una
riga vuota).
*/
func sottostringaMassima(sigma, tau string) {
	str := calcolaSottostringaMassimaComune(sigma, tau)
	fmt.Println(str)
}

/*
indiceCacofonia (σ)
Se non esiste alcun mattoncino di nome σ oppure se il mattoncino di nome σ non appartiene ad
alcuna fila, non compie alcuna operazione.
Altrimenti stampa l’indice di cacofonia della fila cui appartiene il mattoncino di nome σ.
*/
func calcolaSottostringaMassimaComune(s1, s2 string) string {
	len1 := len(s1)
	len2 := len(s2)
	matrice := make([][]int, len1+1)
	for i := range matrice {
		matrice[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s1[i-1] == s2[j-1] {
				matrice[i][j] = matrice[i-1][j-1] + 1
			} else {
				matrice[i][j] = max(matrice[i-1][j], matrice[i][j-1])
			}
		}
	}
	var result string
	i, j := len1, len2
	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			result = string(s1[i-1]) + result
			i--
			j--
		} else if matrice[i-1][j] > matrice[i][j-1] {
			i--
		} else {
			j--
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func indiceCacofonia(g gioco, sigma string) {
	for _, f := range *g.FileDisposte {
		if contiene(f, sigma) {
			cacofonia := 0
			for i := 0; i < len(f.Mattoncini)-1; i++ {
				mattoncino1 := f.Mattoncini[i]
				mattoncino2 := f.Mattoncini[i+1]
				fmt.Println("Controllo ", mattoncino1.Sigma, mattoncino2.Sigma)
				strMassimaComune := calcolaSottostringaMassimaComune(mattoncino1.Sigma, mattoncino2.Sigma)
				cacofonia += len(strMassimaComune)
			}

			fmt.Println("Indice di cacofonia: ", cacofonia)
			return
		}
	}
}

/*
costo (σ, α1, α2, . . . , αn)
Se σ non fa parte di alcuna fila, non compie alcuna operazione.
Altrimenti, detta F la fila cui appartiene il mattoncino di nome σ, stampa il costo del passaggio
da F alla sequenza di forme s = α1, α2, . . . , αn (stampa indefinito se il costo di passaggio `e
indefinito).
*/
func costo() {

}
