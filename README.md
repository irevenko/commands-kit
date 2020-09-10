# Commands Kit 📜
First Go project using [cobra](https://github.com/spf13/cobra)<br>
There are some commands which i've implemented using Go
# List of commands 📋
* [myip](#myip- "Goto #myip-🌐") 
* [compare](#compare- "Goto #") 
* [encrypt](#encrypt- "Goto #") 
* [decrypt](#decrypt- "Goto #") 

# myip 🌐
Check your IP settings with this command<br>
Basically it sends GET request to ipinfo.io

# Preview 🔍
<img src="https://i.imgur.com/mhLCXCG.jpg"><br />

# Installation 🔨
For Windows (.exe) [GoogleDrive](https://drive.google.com/file/d/1Du9M463piig79o05puKJnyYqTb2IhE1z/view?usp=sharing) <br>
For Linux (excutable file) [GoogleDrive](https://drive.google.com/file/d/1VMXo_0oJOlrFljkRaQXIxsF5A4rjPma_/view?usp=sharing)

# Windows 🟦
I've found Powershell Installation the best<br>
[Powershell Path and Aliases Tutorial](https://www.youtube.com/watch?v=4e4lGUVRKFs)

```Download (or take one after compiling on your PC) myip.exe and save it somewhere```<br>
```Go to Powershell and check your profile (if profile doesn't exist - create one)```<br>
```test-path $profile``` <br>
```$profile```<br>

```Add the alias to your profile```<br>
```notepad $profile```<br>
```set-location c:\```<br>
```new-alias myip \your_path\myip.exe```<br>
```from powershell you can just type:  myip```<br>
```from cmd type: powershell myip```

# *nix 🐧🍏
```Download (or take one after compiling on your PC) myip excutable file and save it somewhere``` <br />
```Add the alias to your profile``` <br />

```sudo nano ~/.bashrc```<br />
```alias myip='/your_path/./myip'```

# What I Learned 🧠
* How to send GET request
* Error handling in Go

# ToDo
* Update compare command

# License 📑 
(c) 2020 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)
