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
	Mattoncini    []*mattoncino
	BordiSinistri map[string]string
	BordiDestri   map[string]string
}

type scatola struct {
	Mattoncini map[string]*mattoncino
}

type gioco struct {
	Scatola      scatola
	FileDisposte *[]fila
}

func main() {
	scatola := scatola{
		Mattoncini: make(map[string]*mattoncino),
	}
	gioco := gioco{
		Scatola:      scatola,
		FileDisposte: &[]fila{},
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserisci pure: ") //TODO: da rimuovere
	for scanner.Scan() {
		linea := scanner.Text()
		elementi := strings.Fields(linea)

		switch elementi[0] {
		case "m":
			alpha := elementi[1]
			beta := elementi[2]
			sigma := elementi[3]
			inserisciMattoncino(gioco, alpha, beta, sigma)
			break
		case "s":
			sigma := elementi[1]
			stampaMattoncino(gioco, sigma)
			break
		case "d":
			disponiFila(gioco, linea[2:])
			break
		case "S":
			stampaFila(gioco, linea[2:])
			break
		case "e":
			break
		case "f":
			break
		case "M":
			break
		case "i":
			break
		case "c":
			break
		case "q":
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
	if standard {
		f.BordiSinistri[m.Sigma] = m.Alpha
		f.BordiDestri[m.Sigma] = m.Beta
	} else {
		f.BordiSinistri[m.Sigma] = m.Beta
		f.BordiDestri[m.Sigma] = m.Alpha
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

	fila := &fila{Mattoncini: make([]*mattoncino, 0), BordiSinistri: make(map[string]string), BordiDestri: make(map[string]string)}
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
		if mattonciniPresi[nomi[i][1:]] == false {
			mattoncino1 = g.Scatola.Mattoncini[nomi[i][1:]]
			mattonciniPresi[mattoncino1.Sigma] = true
		}
		if mattoncino1 == nil {
			return
		}
		mattonciniPresi[mattoncino1.Sigma] = true
		if mattonciniPresi[nomi[i+1][1:]] == false {
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
		inserisciMattoncinoInFila(fila, mattoncino1, nomi[i][0] == '+')
	}
	var m *mattoncino
	if mattonciniPresi[nomi[len(nomi)-1][1:]] == false {
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
	if f.BordiDestri[sigma] == "" {
		return false
	}
	return true
}

func stampaFila(g gioco, sigma string) {
	for _, f := range *g.FileDisposte {
		if contiene(f, sigma) {
			fmt.Println("(")
			for _, elem := range f.Mattoncini {
				fmt.Printf("%s: %s, %s\n", elem.Sigma, elem.Alpha, elem.Beta)
			}
			fmt.Println(")")
			return
		}
	}
}

/*
eliminaFila (σ)
Se non esiste alcun mattoncino di nome σ, oppure se il mattoncino di nome σ non appartiene ad
alcuna fila sul tavolo da gioco, non compie alcuna operazione. Altrimenti, sia F la fila cui appartiene
il mattoncino di nome σ. La fila F `e rimossa dal tavolo e tutti i mattoncini che la compongono
sono rimessi nella scatola
*/
func eliminaFila() {

}

/*
disponiFilaMinima (α, β)
Crea e posiziona sul tavolo da gioco una fila di lunghezza minima da α a β. Tutti i mattoncini
della fila devono essere presi dalla scatola. Se non `e possibile creare alcuna fila da α a β, stampa il
messaggio: non esiste fila da α a β
*/
func disponiFilaMinima() {

}

/*
sottostringaMassima (σ, τ )
Stampa su una nuova riga una sottostringa massima ρ di σ e τ (se ρ `e la stringa nulla, stampa una
riga vuota).
*/
func sottostringaMassima() {

}

/*
indiceCacofonia (σ)
Se non esiste alcun mattoncino di nome σ oppure se il mattoncino di nome σ non appartiene ad
alcuna fila, non compie alcuna operazione.
Altrimenti stampa l’indice di cacofonia della fila cui appartiene il mattoncino di nome σ.
*/
func indiceCacofonia() {

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
