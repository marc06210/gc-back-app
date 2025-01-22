# Status

Pour l'instant nous avons un simple endpoint qui liste les publications en se
connectant à une base de donnée postgres.

# A faire

## HTTP
* sécuriser certains endpoints

## BDD
* rendre paramétrable les informations de connexion à la database => done
* tests unitaire bdd ??
* faire de la pagination
* exposer un POST pour ajouter une ligne
* exposer un endpoint pour modifier une ligne (doit être sécurisée)

## Mongodb
* Ajouter la connexion à une base MongoDB et mettre la recherche (cf ancien projet)
* Gérer les articles internes
