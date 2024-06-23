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

**🔧System Informations**

```
🦅 Keagle Stealer - System Informations

👤 Username: Gabriel
🔬 Desktop Name: DESKTOP-TE3EV21
📺 Operating System: Microsoft Windows 10 Home
🔧 HWID: 9bc5e380-e753-46ed-a976-027bd990d5db
⚙️ Processor: Intel(R) Core(TM) i7-10700F CPU @ 2.90GHz
🎞 Memory: 32GB
💾 Disk: 28GB

🔑 Windows Key: 47JN2-2WF7D-9YTKM-VMF8Q-F37RH
``` 

- **Explication**:
  
    - 👤 Username: Indique le nom d'utilisateur actuel de la session Windows. Ici, l'utilisateur est "Gabriel".
    - 🔬 Desktop Name: Montre le nom attribué à la machine dans le réseau ou localement. Dans cet exemple, le nom est "DESKTOP-TE3EV21".
    - 📺 Operating System: Affiche le système d'exploitation installé sur la machine. Ici, il s'agit de "Microsoft Windows 10 Home".
    - 🔧 HWID: Représente l'identifiant matériel unique de la machine (Hardware ID). C'est un identifiant unique pour la machine, ici "9bc5e380-e753-46ed-a976-027bd990d5db".
    - ⚙️ Processor: Indique le type et le modèle du processeur installé. Dans cet exemple, c'est un "Intel(R) Core(TM) i7-10700F CPU @ 2.90GHz".
    - 🎞 Memory: Montre la quantité de mémoire vive (RAM) disponible sur la machine. Ici, la machine dispose de 32GB de RAM.
    - 💾 Disk: Indique l'espace disque disponible sur la machine. Dans cet exemple, il y a 28GB d'espace disque disponible.
    - 🔑 Windows Key: Affiche la clé de produit Windows utilisée pour activer le système d'exploitation. Ici, la clé est "47JN2-2WF7D-9YTKM-VMF8Q-F37RH".


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
