# Groupie Tracker 

## Description 
Groupie Tracker, is an online web application that is taking data from a given API and presents them in the order given on the port `:8080` of the `localhost`.

## API content and app features
The data of the API are basically informations about some famous music bands or solo music artists, that the user can see in the app with ease.
Some basic information examples are :
- **Locations :** some of the locations that the band or artist has played before 
- **Dates :** some of the dates of their music concert
- **Relations :** a combination of the locations and dates that are matching 

but you can see a lot more.

## Usage
To run this program you need to install **golang** and after :
1. Clone the repository :
- Open a terminal, and type the following command :
    ```bash
    git clone https://platform.zone01.gr/git/kpetrout/groupie-tracker.git
2. Run the program :
- You have first get into the repository you just cloned and then run the command:
    ```go
    go run main.go 
- If the program run succesfully you will see a message :
    ```
    Server is running on http://localhost:8080
3. Open in browser :
- Now you can (ctrl+click) on the link of the message or just click here http://localhost:8080

## Authors
Creators and Primary Developers :
- Konstantinos Petroutsos
- Christos Markos
- Christos Gkaldanidis

# Songs that need to be louder
- 3 Pink Floyd
- 5 xxxtentacion ?
- 6 Mac Miller ?
- 12 Rihanna ?
- 13 Genesis
- 14 Phil Collins
- 15 Led Zeppelin
- 21 Mamonas Assassinas
- 46 Coldplay ?

# Locations to change
- georgia usa
- massachusetts usa,
- alabama usa

# Things to be done
**Front-End**
- Update welcomepage
- Suggestions
- Update error pages **(Almost done)**
- Update interface when there are no artists **(Almost done)**
- Update details on filters front-end
- Filters asynchronous **(optional)**
- Filters Locations
- Update the whole details front-end