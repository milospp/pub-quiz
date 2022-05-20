# pub-quiz
Web aplikacija za organizovanje pub kvizova

Autor: Miloš Popović SW-24-2018


## Opis aplikacije
Glavna funkcionalnost aplikacije je organizovanje kvizova, igrači (timovi) pristupaju sobi kviza i dobijaju sinhronizovano pitanja na koja odgovaraju u vremenskom intervalu. Kviz čine:
- **Organizator** koji kontroliše tok igre
- **Igrači** koji daju odgovore u predviđenom obliku (Radio button, check box, input polje...)
- **Nadgledači**, koristi se za projektorski prikaz pitanja, tabele rezultata i ostalih javnih informacija tokom igre (kojima upravlja organizator)

## Funkcionalni zahtevi
Podržati 4 vrste koristnika
- **Administrator**\
Dodaje pitanja za kvizove u aplikaciju (po kategorijama, težini)\
Uvid u poršle buduće i trenutne igre (Filtriranje i Sortiranje)\
Odobrava, menja i briše turnire

- **Neautentifikovani korisnici**\
Registracija, unosi se Ime, Prezime, e-mail
Ulazak u sobu
Kreiranje privatnog turnira (quick)
Pregled javnih turnira
...

- **Registrovani korisnici**\
Izmena profila
Organizovanje javnog turnira\
Prijava/odjava za učestovanje na turniru\
Ulazak u sobu za kviz\
Pravljenje sobe za kviz\

## Scenario
Neregistrovan ili Registrovan korinik napravi privatnu sobu
Izabere pitanja, doda neka svoja, (organizuje tok igre)
Može da postavi keypass dodatno
Dobija kôd od sobe za pristup
Ostali korisnici mogu da uđu u sobu (ne moraju biti registrovani)

Organizator pokrene igru, svi korisnici dobijaju pitanje i neki od načina za unos odgovora u zavisnosti od tipa pitanja
Organizator na svom uređaju ima informacije o tome ko je i kada odgovorio, ima opciju da otkrije tačan odgovor, nakon toga da prikaže sledeće pitanje, da prikaže ponudjene odgovore i tako do kraja kviza.

Na kraju kviza moguće je pregledati izveštaje ko je dao najviše tačnih, netačnih odgovora, ko je najbrži bio...

## Arhitektura

Korisnički servis - Go
Servis za kreiranje kvizova - Go
Servis za vođenje kviza (slanje odgovora, soketi) - Go
Servis za izveštaje kviza - Rust
Web aplikacija - React

Baza
Podaci se čuvaju u PostgreSQL bazi


