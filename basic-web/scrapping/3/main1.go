package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var body = strings.NewReader(`                           
<html>

<head>
<LINK rel='stylesheet' type='text/css' href='bootstrapc.css'>
<link rel=stylesheet type='text/css' href='/css/font-awesome.min.css'>
</head>

<body bgcolor=#ffffff  link=#0000aa vlink=#0000aa>
<font face='arial,helvetica' size=2>
   
   <div class='page-header ke-kanan'><h2><i class='fa fa-cogs'></i> Jadwal Kegiatan Mengajar </h2>
	<div class='sub-headerv2'></div>
	<div class='sub-headerv2'>SEMESTER GASAL 2019 / 2020</div>
   </div>
   <a href=/carijadwal.html><i class='fa fa-reply'></i></a>
   <br/>
   <br/>
   <table border=0 bgcolor=#ffffff cellpadding=4 cellspacing=0 class='tabel-form'>
   <tr class='header hijau'>
   <td>Program Studi</td>
   <td>Jurusan</td>
   <td>Kode MK</td>
   <td>Nama MK</td>
   <td>Sks</td>
   <td>Kls</td>
   <td>Jml Mhs</td>
   <td>Hari</td>
   <td>Jam</td>
   <td>Ruang</td>
   <td>Pasangan</td>
   </tr>
   <tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000101</td>
      <td>Olah Raga I</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-08:20&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Prijoto, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232163</td>
      <td>Struktur Data</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Herry Sofyan, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240343</td>
      <td>Sistem Informasi Geografi</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>SI : PATT. III - 3A&nbsp;</td>
      <td>Budi Santosa, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232182</td>
      <td>Komputasi Numerik</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Heru Cahya Rustamaji, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215162</td>
      <td>Termodinamika</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-3A&nbsp;</td>
      <td>ZUBAIDI ACHMAD, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215032</td>
      <td>Kimia Analisa</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>TITIK MAHARGIANI, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232703</td>
      <td>Komunikasi Data Nirkabel</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220322</td>
      <td>Manajemen Produksi</td>
      <td>2</td>
      <td>A</td>
      <td>15</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220352</td>
      <td>Metode Heuristik</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000101</td>
      <td>Olah Raga I</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-08:20&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Hanafi Mustofa, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220413</td>
      <td>Perancangan Tata Letak Fasilitas</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Intan Berlianty, ST.MT<br>Yuli Dwi Astanti, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240182</td>
      <td>Teori dan Perilaku Organisasi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>SI : PATT. III - 3B&nbsp;</td>
      <td>Hari Prapcoyo, S.Kom., M.ICT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000101</td>
      <td>Olah Raga I</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-08:20&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000101</td>
      <td>Olah Raga I</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-08:20&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Prijoto, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000101</td>
      <td>Olah Raga I</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-08:20&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Hanafi Mustofa, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215422</td>
      <td>Teknik Reaksi Kimia</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>ENDANG SULISTYOWATI, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000101</td>
      <td>Olah Raga 1</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Dr. Lilik Indriharta, S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000101</td>
      <td>Olah Raga 1</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Hanafi Mustofa, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000101</td>
      <td>Olah Raga 1</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000101</td>
      <td>Olah Raga 1</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Sumintarsih, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000101</td>
      <td>Olah Raga 1</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>LAP SEPAK BOLA CC&nbsp;</td>
      <td>Tri Saptono, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232382</td>
      <td>Riset Operasi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220312</td>
      <td>Sistem Pengembangan Produk</td>
      <td>2</td>
      <td>A</td>
      <td>42</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220183</td>
      <td>Pengendalian dan Penjaminan Mutu</td>
      <td>3</td>
      <td>E</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220193</td>
      <td>Proses Manufaktur</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220163</td>
      <td>Penelitian Operasional I</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Puryani, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232332</td>
      <td>Kriptografi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232462</td>
      <td>Uji Kualitas Perangkat Lunak</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232422</td>
      <td>Metodologi Penulisan Ilmiah</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220502</td>
      <td>Ergo. Industri & Interaksi Man. Mesin (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>2</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220152</td>
      <td>Elemen Mesin</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240191</td>
      <td>Praktikum Struktur Data</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>K</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232501</td>
      <td>Praktikum Teknologi Cloud Computing</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>L</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232401</td>
      <td>Praktikum Internet of Things (IoT)</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215062</td>
      <td>Praktikum Kimia Analisa</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215052</td>
      <td>Praktikum Fisika </td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215252</td>
      <td>Praktikum Utilitas</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000101</td>
      <td>Olahraga I</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:30-09:20&nbsp;</td>
      <td>LAP SPK BOLA CC (SI)&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000101</td>
      <td>Olahraga I</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>08:30-09:20&nbsp;</td>
      <td>LAP SPK BOLA CC (SI)&nbsp;</td>
      <td>Dr. Lilik Indriharta, S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220393</td>
      <td>Simulasi Sistem Industri</td>
      <td>3</td>
      <td>A</td>
      <td>16</td>
      <td>Senin&nbsp;</td>
      <td>09:30-12:00&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215182</td>
      <td>Ilmu Bahan Dan Korosi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. III-1B&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215442</td>
      <td>Perpindahan Kalor</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>ENDANG SULISTYOWATI, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215012</td>
      <td>Fisika</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>WASIR NURI, Ir..MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220342</td>
      <td>Kesehatan dan Keselamatan Kerja</td>
      <td>2</td>
      <td>A</td>
      <td>22</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220332</td>
      <td>Metodologi Penelitian Teknik Industri</td>
      <td>2</td>
      <td>A</td>
      <td>20</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215192</td>
      <td>Teknologi Bioproses</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220172</td>
      <td>Analisis Biaya</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Trismi Ristyowati, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232022</td>
      <td>Kalkulus</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232133</td>
      <td>Sistem dan Teknologi Informasi</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Yuli Fauziah, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215052</td>
      <td>Praktikum Fisika </td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:00&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215062</td>
      <td>Praktikum Kimia Analisa</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:00&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Mitha Puspitasari, ST., M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220073</td>
      <td>Pengantar Teknik Industri</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Rifki Indra P, S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232052</td>
      <td>Logika Informatika</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232462</td>
      <td>Uji Kualitas Perangkat Lunak</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240023</td>
      <td>Algoritma dan Pemrograman I</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>SI : PATT. III - 3A&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>F</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Budi Santosa, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232382</td>
      <td>Riset Operasi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220152</td>
      <td>Elemen Mesin</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:00&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232593</td>
      <td>Pemetaan Bawah Permukaan</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Frans Richard Kodong, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>J</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232501</td>
      <td>Praktikum Teknologi Cloud Computing</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240201</td>
      <td>Praktikum Pemrograman Berorientasi Objek</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. DIGITAL (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240143</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>SI : PATT. II - 3C&nbsp;</td>
      <td>Heru Cahya Rustamaji, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240023</td>
      <td>Algoritma dan Pemrograman I</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232052</td>
      <td>Logika Informatika</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>J</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1200022</td>
      <td>Technopreneurship</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. I - 3D&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232022</td>
      <td>Kalkulus</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232163</td>
      <td>Struktur Data</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Herry Sofyan, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>G</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232192</td>
      <td>Interaksi Manusia dan Komputer</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Dyah Ayu Irawati, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232482</td>
      <td>Teknologi Cloud Computing</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220452</td>
      <td>Psikologi Industri</td>
      <td>2</td>
      <td>A</td>
      <td>31</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220312</td>
      <td>Sistem Pengembangan Produk</td>
      <td>2</td>
      <td>B</td>
      <td>20</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220172</td>
      <td>Analisis Biaya</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Trismi Ristyowati, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:40&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220193</td>
      <td>Proses Manufaktur</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200132</td>
      <td>Pemrograman Komputer</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000132</td>
      <td>Bahasa Inggris</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>J</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1200012</td>
      <td>Technopreneurship</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-14:40&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240201</td>
      <td>Praktikum Pemrograman Berorientasi Objek</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. DIGITAL (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220393</td>
      <td>Simulasi Sistem Industri</td>
      <td>3</td>
      <td>B</td>
      <td>41</td>
      <td>Senin&nbsp;</td>
      <td>14:30-17:00&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220413</td>
      <td>Perancangan Tata Letak Fasilitas</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>14:30-17:00&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Intan Berlianty, ST.MT<br>Yuli Dwi Astanti, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>J</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232783</td>
      <td>Pembelajaran Mesin / Machine Learning</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Anggit Ferdita Nugraha, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232362</td>
      <td>Internet of Things (IoT)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Dr. Novrido Charibaldi, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220022</td>
      <td>Kimia Dasar</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200132</td>
      <td>Pemrograman Komputer</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>H</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240293</td>
      <td>Rancang Bangun Perangkat Lunak</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>SI : PATT. II - 3C&nbsp;</td>
      <td>Simon Pulung Nugroho, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232303</td>
      <td>Pemrograman Berorientasi Objek</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Andiko Putro Suryotomo, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232192</td>
      <td>Interaksi Manusia dan Komputer</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Dyah Ayu Irawati, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232013</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Rifki Indra P, S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>SI : PATT. I - 3D&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232653</td>
      <td>Interoperabilitas</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215222</td>
      <td>Praktikum Ilmu Bahan dan Korosi</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220052</td>
      <td>Menggambar Teknik</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Senin&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232352</td>
      <td>Kapita Selekta</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220073</td>
      <td>Pengantar Teknik Industri</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200022</td>
      <td>Technopreneurship</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220612</td>
      <td>Desain Eksperimen (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>15</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220163</td>
      <td>Penelitian Operasional I</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Puryani, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200132</td>
      <td>Pemrograman Komputer</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240323</td>
      <td>Manajemen Investasi TI</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Riza Prapascatama Agusdin, S.Kom., MJM.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232032</td>
      <td>Algoritma dan Pemrograman </td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232733</td>
      <td>Pengenalan Pola</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Heru Cahya Rustamaji, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232163</td>
      <td>Struktur Data</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Herry Sofyan, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220432</td>
      <td>Sistem Rantai Pasok</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232202</td>
      <td>Sistem Digital</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220532</td>
      <td>Perencanaan Lingkungan Kerja (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220322</td>
      <td>Manajemen Produksi</td>
      <td>2</td>
      <td>B</td>
      <td>37</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240162</td>
      <td>Sistem Operasi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>SI : PATT. III - 3B&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215472</td>
      <td>Praktikum Pemisahan Mekanik</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220104</td>
      <td>Matematika Optimasi</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215022</td>
      <td>Matematika</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215172</td>
      <td>Pemisahan Difusional </td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>TUNJUNG WAHYU WIDAYATI, Ir., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>K</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240191</td>
      <td>Praktikum Struktur Data</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. GEOINFORTK (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232311</td>
      <td>Praktikum Sistem Cerdas dan Pendukung Keputusan</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215012</td>
      <td>Fisika</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>WASIR NURI, Ir..MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220104</td>
      <td>Matematika Optimasi</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215202</td>
      <td>Utilitas</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>FAIZAH HADI, Hj.,Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220362</td>
      <td>Sistem Industri Berkelanjutan</td>
      <td>2</td>
      <td>A</td>
      <td>45</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000082</td>
      <td>Pendidikan Kewarganegaraan</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220432</td>
      <td>Sistem Rantai Pasok</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220092</td>
      <td>Mekanika Teknik</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200132</td>
      <td>Pemrograman Komputer</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232032</td>
      <td>Algoritma dan Pemrograman </td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220052</td>
      <td>Menggambar Teknik</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232013</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Rifki Indra P, S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232352</td>
      <td>Kapita Selekta</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232182</td>
      <td>Komputasi Numerik</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Heru Cahya Rustamaji, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232202</td>
      <td>Sistem Digital</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215162</td>
      <td>Termodinamika</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>D3TK : Patt. III-1B&nbsp;</td>
      <td>ZUBAIDI ACHMAD, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215432</td>
      <td>Pemisahan Mekanik</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232133</td>
      <td>Sistem dan Teknologi Informasi</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Yuli Fauziah, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240133</td>
      <td>Struktur Data</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Herry Sofyan, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240392</td>
      <td>Data Mining dan Data Warehousing</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>SI : PATT. III - 3B&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240051</td>
      <td>Praktikum Manajemen Basis Data</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. MULTIMEDIA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>L</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240051</td>
      <td>Praktikum Manajemen Basis Data</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. BASIS DATA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232311</td>
      <td>Praktikum Sistem Cerdas dan Pendukung Keputusan</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240162</td>
      <td>Sistem Operasi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. III - 3C&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220381</td>
      <td>Praktikum Analisis & Perancangan Sistem Informasi</td>
      <td>1</td>
      <td>E</td>
      <td>2</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220073</td>
      <td>Pengantar Teknik Industri</td>
      <td>3</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220163</td>
      <td>Penelitian Operasional I</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Puryani, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215212</td>
      <td>Praktikum Termodinamika</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Budi Santosa, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215182</td>
      <td>Ilmu Bahan Dan Korosi</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-13:40&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200022</td>
      <td>Technopreneurship</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232332</td>
      <td>Kriptografi</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232022</td>
      <td>Kalkulus</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220352</td>
      <td>Metode Heuristik</td>
      <td>2</td>
      <td>C</td>
      <td>35</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232213</td>
      <td>Geoinformatika</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Frans Richard Kodong, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000082</td>
      <td>Pendidikan Kewarganegaraan</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220114</td>
      <td>Statistika Industri</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Trismi Ristyowati, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240312</td>
      <td>Manajemen Sains</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. I - 3D&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220052</td>
      <td>Menggambar Teknik</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220022</td>
      <td>Kimia Dasar</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>F</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240012</td>
      <td>Pengantar Teknologi Informasi dan Komunikasi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Yuli Fauziah, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240191</td>
      <td>Praktikum Struktur Data</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. GEOINFORTK (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232501</td>
      <td>Praktikum Teknologi Cloud Computing</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220114</td>
      <td>Statistika Industri</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>14:30-16:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Trismi Ristyowati, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215242</td>
      <td>Praktikum Pemisahan Difusional</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>14:30-16:30&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Mitha Puspitasari, ST., M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240182</td>
      <td>Teori dan Perilaku Organisasi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Hari Prapcoyo, S.Kom., M.ICT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232273</td>
      <td>Sistem Cerdas dan Pendukung Keputusan</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220052</td>
      <td>Menggambar Teknik</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232163</td>
      <td>Struktur Data</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232422</td>
      <td>Metodologi Penulisan Ilmiah</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232362</td>
      <td>Internet of Things (IoT)</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Rifki Indra P, S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232213</td>
      <td>Geoinformatika</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Frans Richard Kodong, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232563</td>
      <td>Global Positioning System</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Budi Santosa, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220022</td>
      <td>Kimia Dasar</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220381</td>
      <td>Praktikum Analisis & Perancangan Sistem Informasi</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232172</td>
      <td>Komputer dan Masyarakat</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Rochmat Husaini, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220622</td>
      <td>Analisis Multivariat (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Selasa&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220183</td>
      <td>Pengendalian dan Penjaminan Mutu</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220172</td>
      <td>Analisis Biaya</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Trismi Ristyowati, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220152</td>
      <td>Elemen Mesin</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240133</td>
      <td>Struktur Data</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>SI : PATT. II - 3B&nbsp;</td>
      <td>Herry Sofyan, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232202</td>
      <td>Sistem Digital</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220372</td>
      <td>Analisis dan Peranc. Sistem Informasi</td>
      <td>2</td>
      <td>A</td>
      <td>16</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240033</td>
      <td>Kalkulus</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>SI : PATT. I - 3C&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232052</td>
      <td>Logika Informatika</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240312</td>
      <td>Manajemen Sains</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232283</td>
      <td>Rekayasa Perangkat Lunak</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Simon Pulung Nugroho, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215022</td>
      <td>Matematika</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-3A&nbsp;</td>
      <td>PURWO SUBAGYO, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220052</td>
      <td>Menggambar Teknik</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232122</td>
      <td>Matrik dan Ruang Vektor</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Shoffan Saifullah, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Rifki Indra P, S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232213</td>
      <td>Geoinformatika</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Frans Richard Kodong, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215482</td>
      <td>Praktikum Perpindahan Kalor</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Mitha Puspitasari, ST., M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa </td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Ir ABDULLAH KUNTAARSA, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232462</td>
      <td>Uji Kualitas Perangkat Lunak</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215192</td>
      <td>Teknologi Bioproses</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>TITIK MAHARGIANI, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220463</td>
      <td>Analisis dan Perancangan Perusahaan</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220332</td>
      <td>Metodologi Penelitian Teknik Industri</td>
      <td>2</td>
      <td>B</td>
      <td>41</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215172</td>
      <td>Pemisahan Difusional </td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>TUTIK MUJI SETYONINGRUM, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>08:00-15:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>M</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240051</td>
      <td>Praktikum Manajemen Basis Data</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. BASIS DATA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240051</td>
      <td>Praktikum Manajemen Basis Data</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215192</td>
      <td>Teknologi Bioproses</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>K</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220172</td>
      <td>Analisis Biaya</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Trismi Ristyowati, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa </td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>Ir ABDULLAH KUNTAARSA, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220372</td>
      <td>Analisis dan Peranc. Sistem Informasi</td>
      <td>2</td>
      <td>B</td>
      <td>39</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215032</td>
      <td>Kimia Analisa</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>TITIK MAHARGIANI, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220362</td>
      <td>Sistem Industri Berkelanjutan</td>
      <td>2</td>
      <td>B</td>
      <td>17</td>
      <td>Rabu&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Budi Santosa, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232273</td>
      <td>Sistem Cerdas dan Pendukung Keputusan</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232022</td>
      <td>Kalkulus</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232283</td>
      <td>Rekayasa Perangkat Lunak</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Simon Pulung Nugroho, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232032</td>
      <td>Algoritma dan Pemrograman </td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>G</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232043</td>
      <td>Sistem/Teknologi Basis Data</td>
      <td>3</td>
      <td>I</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232133</td>
      <td>Sistem dan Teknologi Informasi</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Yuli Fauziah, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220193</td>
      <td>Proses Manufaktur</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>TI-I : PAT. II-2D&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220732</td>
      <td>Manajemen Teknologi (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220512</td>
      <td>Biomekanika Kerja (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215492</td>
      <td>Praktikum Penanganan Limbah Industri</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-12:00&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000082</td>
      <td>Pendidikan Kewarganegaraan</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232501</td>
      <td>Praktikum Teknologi Cloud Computing</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240061</td>
      <td>Praktikum Algoritma Pemrograman I</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. JARINGAN (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240201</td>
      <td>Praktikum Pemrograman Berorientasi Objek</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. BASIS DATA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240351</td>
      <td>Praktikum Sistem Informasi Geografi</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. MULTIMEDIA (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240061</td>
      <td>Praktikum Algoritma Pemrograman I</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. KOMPUTASI (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220352</td>
      <td>Metode Heuristik</td>
      <td>2</td>
      <td>A</td>
      <td>25</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220672</td>
      <td>Sistem Manajemen Lingkungan (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220762</td>
      <td>Sistem Persediaan (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>8</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232813</td>
      <td>Pengolahan Citra</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232182</td>
      <td>Komputasi Numerik</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Heru Cahya Rustamaji, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220381</td>
      <td>Praktikum Analisis & Perancangan Sistem Informasi</td>
      <td>1</td>
      <td>F</td>
      <td>12</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1200022</td>
      <td>Technopreneurship</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. I - 3C&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240173</td>
      <td>Pemrograman Berorintasi Objek</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>SI : PATT. III - 3A&nbsp;</td>
      <td>Dr. Novrido Charibaldi, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240012</td>
      <td>Pengantar Teknologi Informasi dan Komunikasi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. II - 3B&nbsp;</td>
      <td>Yuli Fauziah, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232273</td>
      <td>Sistem Cerdas dan Pendukung Keputusan</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232332</td>
      <td>Kriptografi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Mangaras Yanu F, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215242</td>
      <td>Praktikum Pemisahan Difusional</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000072</td>
      <td>Pendidikan Pancasila</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000082</td>
      <td>Pendidikan Kewarganegaraan</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Wahyu Wibowo Eko Y., S.Pd., M.M.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220073</td>
      <td>Pengantar Teknik Industri</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240061</td>
      <td>Praktikum Algoritma Pemrograman I</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. JARINGAN (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240351</td>
      <td>Praktikum Sistem Informasi Geografi</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. GEOINFORTK (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-14:40&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240061</td>
      <td>Praktikum Algoritma Pemrograman I</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. KOMPUTASI (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>14:30-16:10&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215232</td>
      <td>Praktikum Bioproses</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>14:30-16:30&nbsp;</td>
      <td>LAB. KIMIA ORGANIK&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220183</td>
      <td>Pengendalian dan Penjaminan Mutu</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>14:30-17:00&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220652</td>
      <td>Programa Non Linear (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240503</td>
      <td>Manajemen Pengelolaan Lapangan Migas</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>SI : PATT. II - 3C&nbsp;</td>
      <td>Herry Sofyan, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220722</td>
      <td>Manajemen Proyek (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>15</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000072</td>
      <td>Pendidikan Pancasila</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232763</td>
      <td>Kecerdasan Bisnis / Business Intelligence</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Vynska Amalia Permadi, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220381</td>
      <td>Praktikum Analisis & Perancangan Sistem Informasi</td>
      <td>1</td>
      <td>C</td>
      <td>15</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220702</td>
      <td>Manajemen Energi (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>1</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>J</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220041</td>
      <td>Praktikum Fisika Dasar</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232362</td>
      <td>Internet of Things (IoT)</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Rifki Indra P, S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220522</td>
      <td>Ergonomi Makro (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>1</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Intan Berlianty, ST.MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>SI : PATT. I - 3C&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232213</td>
      <td>Geoinformatika</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Andiko Putro Suryotomo, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220022</td>
      <td>Kimia Dasar</td>
      <td>2</td>
      <td>A</td>
      <td>2</td>
      <td>Rabu&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa </td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-3A&nbsp;</td>
      <td>Ir ABDULLAH KUNTAARSA, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220862</td>
      <td>Sistem Manufaktur Terotomasi (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215162</td>
      <td>Termodinamika</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>ZUBAIDI ACHMAD, Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215032</td>
      <td>Kimia Analisa</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215212</td>
      <td>Praktikum Termodinamika</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Mitha Puspitasari, ST., M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220183</td>
      <td>Pengendalian dan Penjaminan Mutu</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220163</td>
      <td>Penelitian Operasional I</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Puryani, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220073</td>
      <td>Pengantar Teknik Industri</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Intan Berlianty, ST.MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220302</td>
      <td>Pemodelan Sistem</td>
      <td>2</td>
      <td>A</td>
      <td>1</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220632</td>
      <td>Rekayasa Kualitas (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>1</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000122</td>
      <td>Bahasa Indonesia</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Nur Indrianti, M.T., D.Eng., IPM&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232482</td>
      <td>Teknologi Cloud Computing</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Sylvert Prian Tahalea, S.Si., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Dr. Heriyanto, A.Md., S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240143</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>SI : PATT. I - 3D&nbsp;</td>
      <td>Heru Cahya Rustamaji, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1200022</td>
      <td>Technopreneurship</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232013</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Dr. Novrido Charibaldi, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240152</td>
      <td>Statistika Bisnis</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>SI : PATT. I - 3C&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232462</td>
      <td>Uji Kualitas Perangkat Lunak</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. I - 3A&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220381</td>
      <td>Praktikum Analisis & Perancangan Sistem Informasi</td>
      <td>1</td>
      <td>D</td>
      <td>15</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232482</td>
      <td>Teknologi Cloud Computing</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. III - 3B&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232283</td>
      <td>Rekayasa Perangkat Lunak</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Simon Pulung Nugroho, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232133</td>
      <td>Sistem dan Teknologi Informasi</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Yuli Fauziah, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220261</td>
      <td>Praktikum Perencanaan dan Pengendalian Produksi</td>
      <td>1</td>
      <td>E</td>
      <td>11</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>L. SIST. PROD. (TI)&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220421</td>
      <td>Praktikum Peranc. Tata Letak Fasilitas</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Yuli Dwi Astanti, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215012</td>
      <td>Fisika</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-3A&nbsp;</td>
      <td>WASIR NURI, Ir..MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215452</td>
      <td>Penanganan Limbah Industri </td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215202</td>
      <td>Utilitas</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>FAIZAH HADI, Hj.,Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220463</td>
      <td>Analisis dan Perancangan Perusahaan</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>09:30-12:00&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232202</td>
      <td>Sistem Digital</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Dr. Awang Hendrianto P., S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000072</td>
      <td>Pendidikan Pancasila</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000122</td>
      <td>Bahasa Indonesia</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>SI : PATT. III - 3B&nbsp;</td>
      <td>Hermanto, S.Pd., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240043</td>
      <td>Manajemen Basis Data</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>SI : PATT. III - 3A&nbsp;</td>
      <td>Hari Prapcoyo, S.Kom., M.ICT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000132</td>
      <td>Bahasa Inggris</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220193</td>
      <td>Proses Manufaktur</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200012</td>
      <td>Konsep Teknologi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232352</td>
      <td>Kapita Selekta</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3A&nbsp;</td>
      <td>Hidayatulah Himawan, S.T., M.M., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240152</td>
      <td>Statistika Bisnis</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>SI : PATT. II - 3B&nbsp;</td>
      <td>Bambang Yuwono, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220183</td>
      <td>Pengendalian dan Penjaminan Mutu</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>TI-G : PAT. II-2B&nbsp;</td>
      <td>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232032</td>
      <td>Algoritma dan Pemrograman </td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232052</td>
      <td>Logika Informatika</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232553</td>
      <td>Penginderaan Jarak Jauh</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-12:30&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Simon Pulung Nugroho, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220782</td>
      <td>Penjadwalan Produksi (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Puryani, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220552</td>
      <td>Otomasi Proses Produksi (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215222</td>
      <td>Praktikum Ilmu Bahan dan Korosi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:00-12:00&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240361</td>
      <td>Praktikum Rancang Bangun Perangkat Lunak</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. KOMPUTASI (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232311</td>
      <td>Praktikum Sistem Cerdas dan Pendukung Keputusan</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232311</td>
      <td>Praktikum Sistem Cerdas dan Pendukung Keputusan</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>10:30-12:30&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220592</td>
      <td>Logika Fuzzy (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Dr. Agus Ristono, ST.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240043</td>
      <td>Manajemen Basis Data</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Hari Prapcoyo, S.Kom., M.ICT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232222</td>
      <td>Pemrograman Web</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3A&nbsp;</td>
      <td>Dessyanto Boedi P, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220572</td>
      <td>Enterprises Resources Planning (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>15</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ahmad Muhsin, S.T.,M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000122</td>
      <td>Bahasa Indonesia</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>SI : PATT. III - 3A&nbsp;</td>
      <td>Hermanto, S.Pd., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000072</td>
      <td>Pendidikan Pancasila</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. III - 3C&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000132</td>
      <td>Bahasa Inggris</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>TI-H : PAT. II-2C&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232823</td>
      <td>Teknologi Multimedia</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>IF : PATT. III - 3D&nbsp;</td>
      <td>Oliver Samuel S., S.Kom., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220421</td>
      <td>Praktikum Peranc. Tata Letak Fasilitas</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Yuli Dwi Astanti, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232032</td>
      <td>Algoritma dan Pemrograman </td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>Wilis Kaswidjanti, S.Si., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232022</td>
      <td>Kalkulus</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. I - 3A&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240173</td>
      <td>Pemrograman Berorintasi Objek</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-15:00&nbsp;</td>
      <td>SI : PATT. I - 3D&nbsp;</td>
      <td>Dr. Novrido Charibaldi, S.Kom., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232182</td>
      <td>Komputasi Numerik</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Frans Richard Kodong, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215232</td>
      <td>Praktikum Bioproses</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232052</td>
      <td>Logika Informatika</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>12:30-14:10&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Nur Heri Cahyana, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232401</td>
      <td>Praktikum Internet of Things (IoT)</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232241</td>
      <td>Praktikum Pemrograman Web</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240361</td>
      <td>Praktikum Rancang Bangun Perangkat Lunak</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. KOMPUTASI (SI)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232311</td>
      <td>Praktikum Sistem Cerdas dan Pendukung Keputusan</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000082</td>
      <td>Pendidikan Kewarganegaraan</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>13:30-15:10&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Danang Prasetyo, M.Pd.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215242</td>
      <td>Praktikum Pemisahan Difusional</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>14:30-16:30&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220193</td>
      <td>Proses Manufaktur</td>
      <td>3</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>14:30-17:00&nbsp;</td>
      <td>TI-F : PAT. II-1D&nbsp;</td>
      <td>Dr. Sadi, , S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220772</td>
      <td>Perencanaan Fasilitas (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>11</td>
      <td>Kamis&nbsp;</td>
      <td>14:30-15:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Intan Berlianty, ST.MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215252</td>
      <td>Praktikum Utilitas</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>14:30-16:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220463</td>
      <td>Analisis dan Perancangan Perusahaan</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>14:30-17:00&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Eko Nursubiyantoro, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000072</td>
      <td>Pendidikan Pancasila</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3A&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232192</td>
      <td>Interaksi Manusia dan Komputer</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Dyah Ayu Irawati, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220034</td>
      <td>Fisika Dasar</td>
      <td>4</td>
      <td>C</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-E : PAT. II-1C&nbsp;</td>
      <td>M. Shodiq Abdul Khannan, ST., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220061</td>
      <td>Praktikum Menggambar Teknik</td>
      <td>1</td>
      <td>K</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. GAMTEK (TI)&nbsp;</td>
      <td>Tri Wibawa, S.T.,M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220014</td>
      <td>Kalkulus Dasar</td>
      <td>4</td>
      <td>D</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. II - 3D&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000122</td>
      <td>Bahasa Indonesia</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-16:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Hermanto, S.Pd., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220381</td>
      <td>Praktikum Analisis & Perancangan Sistem Informasi</td>
      <td>1</td>
      <td>A</td>
      <td>11</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232663</td>
      <td>Telematika</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Ahmad Taufiq Akbar, S.Si.,M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240303</td>
      <td>Sistem Informasi Akuntansi</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:00-17:30&nbsp;</td>
      <td>SI : PATT. II - 3C&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>1000082</td>
      <td>Pendidikan Kewarganegaraan</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Kamis&nbsp;</td>
      <td>15:30-16:40&nbsp;</td>
      <td>D3TK :  Patt. III-1A&nbsp;</td>
      <td>Danang Prasetyo, M.Pd.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232013</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>D</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:00-09:30&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Agus Sasmito Aribowo, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215212</td>
      <td>Praktikum Termodinamika</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Yuli Ristianingsih, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Tri Saptono, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220742</td>
      <td>Pengukuran Kinerja (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>15</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>IF : PATT. II - 3C&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220752</td>
      <td>Sistem Pemeliharaan (Pil)</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Apriani Soepardi, STP., MT.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232373</td>
      <td>Sistem Operasi</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Bagus Muhammad Akbar, S.ST., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215022</td>
      <td>Matematika</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-3A&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>SI : PATT. I - 3C&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>F</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220152</td>
      <td>Elemen Mesin</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Moch. Chaeron, ST., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220421</td>
      <td>Praktikum Peranc. Tata Letak Fasilitas</td>
      <td>1</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Intan Berlianty, ST.MT<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240033</td>
      <td>Kalkulus</td>
      <td>3</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-10:00&nbsp;</td>
      <td>SI : PATT. II - 3D&nbsp;</td>
      <td>Juwairiah, S.Si., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220261</td>
      <td>Praktikum Perencanaan dan Pengendalian Produksi</td>
      <td>1</td>
      <td>A</td>
      <td>15</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. SIST. PROD. (TI)&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215172</td>
      <td>Pemisahan Difusional </td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>07:30-09:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>TUNJUNG WAHYU WIDAYATI, Ir., MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>N</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232231</td>
      <td>Praktikum Implementasi Struktur Data</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232061</td>
      <td>Praktikum  Algoritma dan Pemrograman</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232501</td>
      <td>Praktikum Teknologi Cloud Computing</td>
      <td>1</td>
      <td>E</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. DIGITAL (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>M</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>08:00-10:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215202</td>
      <td>Utilitas</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-1A&nbsp;</td>
      <td>FAIZAH HADI, Hj.,Ir.,MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232673</td>
      <td>Basis Data Berorientasi Objek</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-12:00&nbsp;</td>
      <td>IF : PATT. I - 3B&nbsp;</td>
      <td>Agus Sasmito Aribowo, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220421</td>
      <td>Praktikum Peranc. Tata Letak Fasilitas</td>
      <td>1</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Yuli Dwi Astanti, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Tri Saptono, S.Pd., M.Or.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215182</td>
      <td>Ilmu Bahan Dan Korosi</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>D3TK : Patt. II-2A&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220261</td>
      <td>Praktikum Perencanaan dan Pengendalian Produksi</td>
      <td>1</td>
      <td>B</td>
      <td>15</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>L. SIST. PROD. (TI)&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220272</td>
      <td>Organisasi dan Manajemen Perusahaan Industri</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>09:30-11:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT<br>Ir. Irwan Soejanto, M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215232</td>
      <td>Praktikum Bioproses</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>10:00-12:00&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>D</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. II - 3B&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1000132</td>
      <td>Bahasa Inggris</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>IF : PATT. I - 3D&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>10:00-11:40&nbsp;</td>
      <td>SI : PATT. I - 3C&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. BASIS DATA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>Q</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. KOMPUTASI (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>O</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. MULTIMEDIA (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>G</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. JARINGAN (IF)&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232071</td>
      <td>Praktikum Implementasi Basis Data</td>
      <td>1</td>
      <td>P</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. GEOINFORMATIKA&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232013</td>
      <td>Matematika Diskrit</td>
      <td>3</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:30&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Agus Sasmito Aribowo, S.Kom., M.Cs.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1240513</td>
      <td>Manajemen Pengelolaan Tambang</td>
      <td>3</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:30&nbsp;</td>
      <td>SI : PATT. I - 3D&nbsp;</td>
      <td>&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215222</td>
      <td>Praktikum Ilmu Bahan dan Korosi</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:00-15:00&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Andri Perdana, S.T., M.Eng.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000012</td>
      <td>Pendidikan Agama Islam</td>
      <td>2</td>
      <td>A</td>
      <td>16</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:10&nbsp;</td>
      <td>TI-C : PAT. I-1D&nbsp;</td>
      <td>Dr. Robby Habiba Abror, SAg, M.Hum&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000132</td>
      <td>Bahasa Inggris</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:10&nbsp;</td>
      <td>TI-D : PAT. II-1B&nbsp;</td>
      <td>Dr. Herlina Jayadianti, S.T., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>H</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-16:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000072</td>
      <td>Pancasila</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:10&nbsp;</td>
      <td>SI : PATT. I - 3B&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1000092</td>
      <td>Bela Negara dan Widya Mwat Yasa</td>
      <td>2</td>
      <td>E</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:10&nbsp;</td>
      <td>TI-A : PAT. I-1B&nbsp;</td>
      <td>Ir. Lagiman, M.Si&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220421</td>
      <td>Praktikum Peranc. Tata Letak Fasilitas</td>
      <td>1</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Yuli Dwi Astanti, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215062</td>
      <td>Praktikum Kimia Analisa</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:30&nbsp;</td>
      <td>LAB. DTK JTK&nbsp;</td>
      <td>ADI ILCHAM, ST.,MT.,DR&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215052</td>
      <td>Praktikum Fisika </td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:30&nbsp;</td>
      <td>LAB OTK UPT L. DASAR&nbsp;</td>
      <td>Susanti Rina Nugraheni, ST,M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220261</td>
      <td>Praktikum Perencanaan dan Pengendalian Produksi</td>
      <td>1</td>
      <td>C</td>
      <td>15</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:30&nbsp;</td>
      <td>L. SIST. PROD. (TI)&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220022</td>
      <td>Kimia Dasar</td>
      <td>2</td>
      <td>C</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>13:30-15:10&nbsp;</td>
      <td>TI-B : PAT. I-1C&nbsp;</td>
      <td>Ir. Dyah Rachmawati Lucitasari, MT&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>Teknik Kimia</td>
      <td>0215252</td>
      <td>Praktikum Utilitas</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>LAB. FISIKA&nbsp;</td>
      <td>Mitha Puspitasari, ST., M.Eng&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>SISTEM INFORMASI</td>
      <td>1000072</td>
      <td>Pancasila</td>
      <td>2</td>
      <td>A</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>15:30-17:10&nbsp;</td>
      <td>SI : PATT. I - 3B&nbsp;</td>
      <td>Yunie Herawati, Ir., M.Hum.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220421</td>
      <td>Praktikum Peranc. Tata Letak Fasilitas</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>15:30-17:30&nbsp;</td>
      <td>LAB. STATISTIK (TI)&nbsp;</td>
      <td>Yuli Dwi Astanti, S.T.,M.T.<br>Berty Dwi Rahmawati, S.T.,M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INFORMATIKA</td>
      <td>1232192</td>
      <td>Interaksi Manusia dan Komputer</td>
      <td>2</td>
      <td>B</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>15:30-17:10&nbsp;</td>
      <td>IF : PATT. I - 3C&nbsp;</td>
      <td>Dyah Ayu Irawati, S.T., M.Kom.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220261</td>
      <td>Praktikum Perencanaan dan Pengendalian Produksi</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Jum'at&nbsp;</td>
      <td>15:30-17:30&nbsp;</td>
      <td>L. SIST. PROD. (TI)&nbsp;</td>
      <td>Laila Nafisah, ST.,MT.<br>Ismianti, S.T., M.Sc&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1220201</td>
      <td>Praktikum Proses Manufaktur</td>
      <td>1</td>
      <td>D</td>
      <td>0</td>
      <td>Sabtu&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. PROS. PROD. (TI)&nbsp;</td>
      <td>Gunawan Madyono Putro, S.T., M.T.<br>Hasan Mastrisiswadi, S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>I</td>
      <td>0</td>
      <td>Sabtu&nbsp;</td>
      <td>07:30-09:30&nbsp;</td>
      <td>L. APK & ERGRO. (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>J</td>
      <td>0</td>
      <td>Sabtu&nbsp;</td>
      <td>09:30-11:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>K</td>
      <td>0</td>
      <td>Sabtu&nbsp;</td>
      <td>12:30-14:30&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Muhk. Nasir Ramdhani, S.M., M.T.&nbsp;</td>
      </tr><tr class='cell'>
      <td></td>
      <td>TEKNIK INDUSTRI</td>
      <td>1200141</td>
      <td>Praktikum Pemrograman Komputer</td>
      <td>1</td>
      <td>L</td>
      <td>0</td>
      <td>Sabtu&nbsp;</td>
      <td>15:00-17:00&nbsp;</td>
      <td>L. KOMPUTASI (TI)&nbsp;</td>
      <td>Sutrisno, S.Si., M.T.<br>Astrid Wahyu A.W., S.T., M.Sc.&nbsp;</td>
      </tr></table></body></html>`)

func main() {
	z := html.NewTokenizer(body)
	content := [][]string{}
	innerContent := []string{}
	counter := 0

	// While have not hit the </html> tag
	for z.Token().Data != "html" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()

			if t.Data == "tr" {
				innerContent = []string{}
				counter = 0
			}

			if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					t := strings.TrimSpace(text)
					innerContent = append(innerContent, t)
					counter++
				}

				// change counter according to your coulumn size
				if counter == 8 {
					content = append(content, innerContent)
				}

			}

		}
	}
	// Print to check the slice's content
	// fmt.Println(content)

	for _, data := range content {
		if data[0] == "TEKNIK INFORMATIKA" {
			fmt.Println(data)
		}
	}
}
