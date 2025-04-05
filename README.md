## DnDSim
A web app that connects to a custom (fine-tuned) ChatGPT API to play the Dungeonmaster in your virtual Pen and Paper game.  
This is currently more of an educational project for myself.
The goal is to offer a more rule-based and transparent alternative to services like [AI Dungeon](https://aidungeon.com/).

### How it works
The web app will keep track of your character sheet (skills, inventory) to send to the API for each request, and the ChatGPT API will adhere to the rules of a popular Pen-and-Paper roleplaying format to show you what dice it will roll, what the outcomes of your attempt may be. This "Dungeonmaster" will also know the scenario and the world you live in beforehand, so the story is less likely to get out of control, and you (the player) can't gaslight the AI as easily into getting what you want.  
The goal is to have a harder, but also fairer and more realistic AI roleplaying experience.  
Currently, the Dungeonmaster API is not available publicly, so you will need to provide your own link to a fine-tuned ChatGPT API.

### Development
This project uses [air](https://github.com/air-verse/air) and [templ](https://github.com/a-h/templ). Both should be installed on your system if you want to contribute to the development or test the project for yourself (It is currently not hosted anywhere). Also, the [TailwindCSS CLI](https://tailwindcss.com/docs/installation) must be installed on your system. Running it or templ youself is not necessary if you use 
```sh
air
```
to start the service.
For that to work as intended, install both
```sh
go install github.com/air-verse/air@latest
```
and 
```sh
go install github.com/a-h/templ/cmd/templ@latest
```

### Technology
This project mainly exists to see how far I can get with a tech stack consisting of Go, Templ, HTMX, and some TailwindCSS for styling and SQLite for persistence.