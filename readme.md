[![CircleCI](https://circleci.com/gh/TanklesXL/FragCollector.svg?style=svg)](https://circleci.com/gh/TanklesXL/FragCollector)

# FragCollector: 

A simple and convenient CLI written in Go for managing your fragrance collection. 

This application was built using Cobra CLI.

Because all fragrance information is scraped from the directory on www.basenotes.net/fragrancedirectory, adding new fragrances will require an internet collection, although viewing/filtering your existing collection does not.

## Quick Start Guide:

### If you have Go installed, go get and go install the application: 

```
go get github.com/TanklesXL/FragCollector
```

This is as simple as using "cd" to get into the FragCollector folder and calling:

```
go install
```

 ### Adding fragrances to your collection

Fragrances can be added to your collection via their name or the url for their entry in the Basenotes directory. Add requires the --name/-n flag (--brand/-b is also available to improve the search, see below), and url requires the --url/-u flag.

#### To search by name:

To add a fragrance to your collection, you simply need to call the add command with the name of the fragrance you would like to search for. While the command is case-insensitive, try to be as close as possible for a search, and to add the  fragrance you will need to type it's name correctly.

For example, assuming that we want to add Aventus by the fragrance house Creed tp our collection, we would do (***note***: a name with multiple words will need to be surrounded by brackets i.e. "Interlude Man")

```
FragCollector add -n Aventus
```

#### To search by fragrance house and by name

If you would like to be more specific with your search (and thus more likely to get the correct search result), you can use the --name/-n and --brand/-b flags together to make it easier to narrow down the search, this method  searches only through the list of fragrances made by the fragrance house you're asking it to search through. Using the same example as above:

```
FragCollector add -n Aventus -b Creed
```

#### To search by URL

You simply need to use the --url/-u flag followed by the URL to the webpage. To add Aventus using this method you would do the following:

```
FragCollector add -u http://www.basenotes.net/ID26131702.html
```

### Removing fragrances from your collection:

This is done by calling:

```
FragCollector remove
```

You will then be shown your collection and asked to select which fragrance to remove, this is done by typing in the number assigned to the fragrance and hitting enter.

### How to view your collection

There are a few ways to view your collection, they are as follows:

1. In alphabetical order by brand:

   ```
   FragCollector mycollection
   ```

2. In alphabetical order solely by name:

   ```
   FragCollector mycollection alpha
   ```

3. In order by the fragrance notes contained in your collection:

   ```
   FragCollector mycollection notes
   ```

4. In order by fragrance note but with every group expanded (use notes with --list/-l:

   ```
   FragCollector mycollection notes -l
   ```

### Getting a Fragrance's Info

By using FragCollector, you can see some useful information about a fragrance in your collection using the command:

```
FragCollector info
```

This command will show you your collection in alphabetical order by brand (the same as doing *FragCollector mycollection*) and prompt you to select the one you would like to expand by typing in the number corresponding to it.

Just like for viewing your collection as a flat list in alphabetical order, this is also an option when displaying a fragrance's info, for this use the --list/-l flag like so:

```
FragCollector info -l
```

