# <div align="center">HANGMAN-WEB
## SOMMAIRE
- [I. Comment installer le Hangman-Web](#i-comment-installer-le-hangman-web)
- [II. Fonctionnement du Hangman-Web](#ii-fonctionnement-du-hangman-web)
- [III. Comment jouer au Hangman-Web](#iii-comment-jouer-au-hangman-web)
- [IV. Hebergement Serveur Hangman-Web](#iv-hebergement-serveur-hangman-web)
- [V. Easter Egg  ](#v-easter-egg)
- [VI. Si le Hangman-Web ne fonctionne pas](#VI-si-le-hangman-web-ne-fonctionne-pas)

## I. Comment installer le Hangman-Web
Pour installer le Hangman-Web, il faut d'abord cloner le repository sur votre ordinateur. Pour cela, il faut ouvrir un terminal et taper la commande suivante :
```bash
git clone https://ytrack.learn.ynov.com/git/cmaxime/Hangman-Web.git
```
Ensuite, il faut se rendre sur VsCode et ouvrir le dossier hangman-web. Une fois le dossier ouvert, il faut ouvrir un terminal et taper la commande suivante :
```go
go run .\main.go
```
## II. Fonctionnement du Hangman-Web
Le Hangman-Web est un jeu de pendu sur le net. Le but est de trouver le mot caché en devinant les lettres qui le composent. Pour cela, le joueur a droit à 10 essais. Si le joueur trouve le mot, il gagne la partie. Sinon, il perd la partie.
## III. Comment jouer au Hangman-Web 
Pour jouer au Hangman-Web, il faut d'abord lancer le serveur. Ensuite, il faut choisir un pseudo puis un niveau parmi les 3 proposés. Une fois le niveau choisi, le joueur doit cliquer sur jouer puis deviner le mot caché en proposant des lettres ou un mot. 
## IV. Hebergement Serveur Hangman-Web
Nous avons hebergé notre hangman-web sur un serveur à l'aide du gourou des infrastructures voici donc le lien : https://diane.ynov-games.com/
## V. Easter Egg 
Rentrez un pseudo puis jouez et taper mentors
## VI. Si le Hangman-Web ne fonctionne pas
Si le Hangman ne fonctionne pas, il faut d'abord vérifier que vous avez bien installé Go sur votre ordinateur. 

Si Go est bien installé, il faut vérifier que vous avez bien installé le repository du Hangman-Web.

Si le repository est bien installé, il faut vérifier que vous avez bien ouvert le dossier hangman-Web sur VsCode.

Si le Hangman de s'execute pas après ces étapes alors supprimer le dossier hangman-Web et recommencer l'installation du Hangman-Web.

Si le Hangman-Web ne fonctionne toujours pas, il faut contacter les auteurs du Hangman-Web.
 
## <div align="right">Les auteurs du Hangman-Web
<div align="right">SAUTEREAU DU PART Diane  
<div align="right">COLOMBAN-FERNANDEZ Luna  
<div align="right">CHORT Maxime