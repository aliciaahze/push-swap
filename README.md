# Push-Swap

## Présentation

Pour ce projet, j'ai dû créer deux programmes en Go pour trier des nombres en utilisant deux piles (A et B).

Le premier programme, push-swap, doit trouver les instructions pour trier les nombres dans l'ordre croissant.

Le deuxième programme checker, sert à vérifier si les instructions permettent bien de trier les nombres.

### Fonctionnement 

On a 11 instructions possibles pour déplacer les nombres entre les deux piles, les échanger ou les faire tourner.

Pour trier les nombres, j'ai utilisé un tri simple pour les listes de 2,3,4 ou 5 nombres et le tri radix pour les listes plus grandes.

Le tri radix fonctionne en regardant les bits des nombres. J'ai aussi utilisé des rangs pour pouvoir trier les nombres négatifs.

## Organisation des fichiers

- cmd/push-swap/main.go : programme qui génère les instructions de tri.
-cmd/checker/main.go : programme qui vérifie les instructions.
- internal/pushswap/stack.go : structure et création des piles.
- internal/pushswap/operations.go : les 11 opérations autorisées.
- internal/pushswap/parsing.go : lecture et conversion des arguments.
- internal/pushswap/validation.go : détection des doublons.
- internal/pushswap/sort.go : fonctions et algorithmes de tri.
- internal/pushswap/checker.go : application des instructions.
- internal/pushswap/*_test.go : tests unitaires.

## Lancer le projet 

Pour lancer push-swap : go run ./cmd/push-swap "3 2 1"

Pour vérifier les instructions avec checker : go run ./cmd/push-swap "3 2 1" | go run ./cmd/checker "3 2 1"

Vérifier les instructions avec Checker : go run ./cmd/push-swap "3 2 1" | go run ./cmd/checker "3 2 1" 

## Tests

J'ai fait plusieurs tests avec des nombres positifs, négatifs, des listes déjà triées et des arguments invalides.

Pour lancer les tests Go : go test ./...


## Outils utilisés

- Go
- VS Code
- GitHub