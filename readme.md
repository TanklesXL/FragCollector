# FragCollector: 

A simple and convenient CLI written in Go for managing your fragrance collection. 

This application was built using Cobra CLI.

Because all fragrance information is scraped from the directory on www.basenotes.net/fragrancedirectory, adding new fragrances will require an internet collection, although viewing/filtering your existing collection does not.

## Quick Start Guide:

### Step 0 (if you have Go installed) get and install the application: 

```
go get github.com/TanklesXL/FragCollector
```

This is as simple as using "cd" to get into the FragCollector folder and calling:

```
go install
```

 ### Step 1: Adding fragrances to your collection



Fragrances can be added to your collection via their name or the url for their entry in the Basenotes directory. Add requires the --name/-n flag (--brand/-b is also available to improve the search, see below), and url requires the --url/-u flag.

#### To search by name:

To add a fragrance to your collection, you simply need to call the add command with the name of the fragrance you would like to search for. While the command is case-insensitive, try to be as close as possible for a search, and to add the  fragrance you will need to type it's name correctly.

For example, assuming that we want to add Aventus by the fragrance house Creed tp our collection, we would do (***note***: a name with multiple words will need to be surrounded by brackets i.e. "Interlude Man")

```
FragCollector add -n Aventus
```

#### To search by fragrance house and by name

If you would like to be more specific with your search (and thus more likely to get the correct search result), you can use the -