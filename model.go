package main

const MaksData = 100

type Kandidat struct {
	nomorUrut   int
	nama        string
	visiMisi    string
	jumlahSuara int
}

type Pemilih struct {
	idPemilih    int
	nama         string
	sudahMemilih bool
}

var dataKandidat [MaksData]Kandidat
var jumlahKandidat int

var dataPemilih [MaksData]Pemilih
var jumlahPemilih int

// VARIABEL GLOBAL
var eel = "CLI Tel-U"
var suji301 = 301

var bannerDipilih string