package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mattoncino struct {
	Alpha string // bordo α
	Beta  string // bordo β
	Sigma string // nome
}

type fila struct {
	Mattoncini    []*Mattoncino
	BordiSinistri map[string]string
	BordiDestri   map[string]string
}

type scatola struct {
	Mattoncini map[string]*Mattoncino
}

type gioco struct {
	Scatola      scatola
	FileDisposte []*fila
}

func main() {
	scatola := scatola{
		Mattoncini: make(map[string]*Mattoncino),
	}
	gioco := gioco{
		Scatola:      scatola,
		FileDisposte: make([]*fila, 0),
	}

	scanner := bufio.NewScanner(os.Stdin)
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
		mattoncino := &Mattoncino{Alpha: alpha, Beta: beta, Sigma: sigma}
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
func disponiFila(g gioco, listaNomi string) {

}

/*
stampaFila (σ)
Se non esiste alcun mattoncino di nome σ, oppure se il mattoncino di nome σ non appartiene
ad alcuna fila sul tavolo da gioco, non compie alcuna operazione. Altrimenti, stampa la fila cui
appartiene il mattoncino con nome σ, secondo il formato specificato nell’apposita sezione.
*/
func stampaFila() {

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
