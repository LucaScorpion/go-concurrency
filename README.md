# Golangang: Concurrency

Deze opdracht is een minimale opzet voor een TCP server, of hub.
De server opent een poort waarmee meerdere clients kunnen verbinden,
en stuurt berichten door naar alle verbonden clients.

## Prerequisites

- Go 1.21
- Netcat

[Netcat](https://netcat.sourceforge.net) is een tool waarmee je eenvoudig TCP verbindingen kan openen,
dit zullen we gebruiken als client om met de server te praten.

- Voor [Homebrew](https://formulae.brew.sh/formula/netcat#default): `brew install netcat`
- Voor Ubuntu: `apt install netcat`

## Quick Start

Start de server:

```shell
go run ./cmd/server
```

Verbind met de server:

```shell
nc localhost 7000
```

Door iets te typen in Netcat wordt dat over de socket naar de server verzonden,
en berichten die binnenkomen worden geprint naar stdout.

## Opdracht

De code waarmee je moet beginnen staat op de `main` branch.
De oplossingen voor de drie delen staan op de `solution-N` branches (waar `N` het nummer van de opdracht is).

### Deel 1

Het eerste doel is om te zorgen dat meerdere clients met de server kunnen verbinden,
en dat de server alles wat binnenkomt broadcast naar alle clients.
Als dit goed werkt zou de server ook alle inkomende berichten moeten printen (prefixed met "< ").

Let op!
Omdat je hier met goroutines gaat werken is het ook belangrijk dat je - waar nodig - gedeelde memory access synchroniseert.

### Deel 2

Nu berichten van clients goed gebroadcast worden, zou het ook handig zijn als de server zelf berichten kan sturen.
Zorg dat je input van stdin op de server afhandelt, en broadcast dit ook naar alle clients.

### Deel 3

Wanneer een client een bericht naar de server stuurt, krijgt deze het bericht ook zelf weer binnen.
Zorg dat de server bij het broadcasten van een bericht de client waar het vandaan komt overslaat.
