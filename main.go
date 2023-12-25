package main

type Mattoncino struct {
	Alpha string // bordo α
	Beta  string // bordo β
	Sigma string // nome
}

type Fila struct {
	Mattoncini    []Mattoncino
	BordiSinistri map[string]string
	BordiDestri   map[string]string
}

type Scatola struct {
	Mattoncini map[string]Mattoncino
}

type Gioco struct {
	Scatola      Scatola
	FileDisposte []Fila
}

/*
inserisciMattoncino (α, β, σ)
Se esiste già un mattoncino di nome σ oppure se α `e uguale a β, non compie alcuna operazione.
Altrimenti, inserisce nella scatola il mattoncino definito dalla tripla (α, β, σ).
*/
func inserisciMattoncino() {

}

/*
stampaMattoncino (σ)
Se non esiste alcun mattoncino di nome σ non compie alcuna operazione. Altrimenti, stampa il
mattoncino con nome σ, secondo il formato specificato nell’apposita sezione.
*/
func stampaMattoncino() {

}

/*
disponiFila (±σ1, ±σ2, . . . , ±σn)
dove ± indica uno dei due simboli + o −. Verifica se nella scatola ci sono i mattoncini di nome
σ1, σ2, . . . , σn e se la sequenza di mattoncini ±σ1, ±σ2, . . . , ±σn costituisce una fila; in questo caso,
toglie dalla scatola i mattoncini che la comp
*/
func disponiFila() {

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
