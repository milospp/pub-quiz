# pub-quiz
Web aplikacija za organizovanje pub kvizova

Autor: Miloš Popović SW-24-2018


## Opis aplikacije
Glavna funkcionalnost aplikacije je organizovanje kvizova.
Kviz može biti privatni (igrači pristupaju preko linka koji dobijaju od organizatora) ili javni (igrači se prijavljuju na neki od dostupnih kvizova)
Vreme početka javnih kvizova može biti zakazano u budućnosti.

Igrači pristupaju sobi kviza i dobijaju sinhronizovano pitanja na koja odgovaraju u vremenskom intervalu. Uloge u jednoj partiji kviza čine:
- **Organizator** koji kontroliše tok igre
- **Igrači** koji daju odgovore u predviđenom obliku (Radio button, check box, input polje...)
- **Nadgledači**, koristi se za projektorski prikaz pitanja, tabele rezultata i ostalih javnih informacija tokom igre. (Predviđeno za prikaz na projektoru ili TV-u)


## Funkcionalni zahtevi
U celokupnoj aplikaciji postoji
- **Moderator**\
Dodaje pitanja u bazu pitanja za kvizove (po kategorijama, težini)\
Uvid u poršle buduće i trenutne igre (Filtriranje i Sortiranje)\
Pregled određenog kviza - Lista pitanja, rezultati
Odobrava, menja i briše turnire

- **Neautentifikovani korisnici**\
Registracija, unosi se Ime, Prezime, e-mail
Ulazak u sobu privatnog kviza
Kreiranje privatnog turnira
Pregled javnih kvizova


- **Registrovani korisnici**\
Sve funkcije neautentifikovanog korisnika
Izmena profila
Organizovanje javnog kviza\
Prijava/odjava za učestovanje na turniru\
Ulazak u sobu za kviz\

## Scenario
### Kreiranje privatnog kviza
Korisnik bira opciju da napravi privatni turnir. Popunjava formu sa pitanjima i pravilima igre. Kada potvrdi izmene dobija KOD i link za pristup. Ima pregled ko je sve pristupio i može da ih izbaci iz sobe.  
Korisnici ulaze u "sobu" kviza preko linka ili unosom koda sobe u aplikaciji, inače se privatni kviz nigde ne vidi. Zamišljen je za privatnu organizaciju.
Korisnik može da izabere da bude posmatrač, samo će videti pitanja bez mogućnosti da učestvuje. Ova opcija se koristi za prikaz pitanja na projektoru

### Kreiranje javnog kviza
Korisnik bira opciju za javni kviz. Bira vreme, naziv, opis, popunjava formu sa pitanjima i pravilima igre.
Turnir se šalje na odobravanje administratoru, ako odobri postaje vidljiv u aplikaciji za sve korisnike i imaju opciju da se prijave.
Igrači se mogu prijaviti i kao posmatrači

### Tok igre
Kada se kviz započne, igra se odvija u ciklusima od nekoliko koraka za svako pitanje
- Pauza pred pitanje
- Ispis pitanja
- Ispis ponuđenih odgovora ako postoje
- Početak vremena za odgovor
- Prikaz tačnog odgovora
- Opciono presek sa statistikom

Prelaz u sledeću fazu kontroliše organizator (u realnom vremenu)

Na kraju kviza moguće je pregledati izveštaje ko je dao najviše tačnih, netačnih odgovora, ko je najbrži bio...

### Otalo
Svaki korisnik igra za sebe, moguće je postojanje tima na način da koriste svi zajedno jedan uređaj za davanje odgovora

### Turniri
Moderatori mogu organizovati turnire koji se sastoje od nekoliko uvezanih kvizova

## Arhitektura
Korisnički servis - Go
Servis za kreiranje kvizova - Go
Servis za vođenje kviza (slanje odgovora, soketi) - Go
Servis za izveštaje kviza - Rust
Web aplikacija - React

Baza
Podaci se čuvaju u PostgreSQL bazi
