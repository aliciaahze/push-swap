# Push-Swap

## Présentation

Pour ce projet, j'ai dû créer deux programmes en Go pour trier des nombres en utilisant deux piles (A et B).

Le premier programme, push-swap, doit trouver les instructions pour trier les nombres dans l'ordre croissant.

Le deuxième programme checker, sert à vérifier si les instructions permettent bien de trier les nombres.

### Fonctionnement 

On a 11 instructions possibles pour déplacer les nombres entre les deux piles, les échanger ou les faire tourner.

Pour trier les nombres, j'ai utilisé un tri simple pour les listes de 2 ou 3 nombres et le tri radix pour les listes plus grandes.

Le tri radix fonctionne en regardant les bits des nombres. J'ai aussi utilisé des rangs pour pouvoir trier les nombres négatifs.

## Organisation des fichiers

cmd/push-swap : programme de tri
cmd/checker : programme qui vérifie le tri
internal/pushswap : fonctions utilisées par les deux programmes

## Comment lancer le projet ?

Pour lancer push-swap : go run ./cmd/push-swap "3 2 1"

Pour vérifier les instructions avec checker : go run ./cmd/push-swap "3 2 1" | go run ./cmd/checker "3 2 1"


Si les nombres sont bien triés, le checker affiche OK sinon, il affiche KO. Si un argument ou une instruction est invalide, le programme affiche Error.

## Tests

J'ai fait plusieurs tests avec des nombres positifs, négatifs, des listes déjà triées et des arguments invalides.

Pour lancer les tests Go : go test ./...


## Outils utilisés

- Go
- VS Code
- GitHub