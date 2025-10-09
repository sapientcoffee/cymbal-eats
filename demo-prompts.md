# Super User

## IDE - Super User

[ ] Chat (traditional) and Inline interaction
[ ] Custom Rules
[ ] Custom Commands
[ ] Add selection to chat context
[ ] Deploy apps to Cloud Run
[ ] Agent Mode
    [ ] MCP

## CLI -  Super User
[ ] Slash commands
    [ ] OOTB (stats, help, tools)
    [ ] Gemini.md
    [ ] .gemini/commands/
[ ] Extensions
[ ] MCP
[ ] Headless mode
[ ] Settings
[ ] Shell mode
[ ] Checkpointing and chat saves
[ ] \ for carrigae returns
[ ] ctl-r for history





# End to end workflow demo

Open PRD in Google Doc

Ask Gemini app to review the PRD and if it has enough to allow a developer to understand. How can it be imporved.

Do the same prompt in Gemini CLI

```
Whats the context of the ITEM-1 user story in jira
```

```
Does this align with the PRD?
```

```
/plan the implementation of this
```

```
execute plan in a new branch
```

## Open new pane for parrellel documentaiton

```
Review the current documentaion for each service and improve it. I want /diagram:new flows and architectures for each service and also how the services interact with each other. Do this in a new branch.
```

## new pane to write unit test and execute (so can see live correction)

```
create  unit test for @coffee-service and ensure they run.
```


```
/security:analyze @coffee-service/main.go
```

Or if use Snyk use their MCP (TODO fix CLI)
```
Run security code scan for employee-ui application and report issues.
```


/code-review

Commit changes to branch and create a PR

<!-- /plan how to add a new api endpoint that will provide a productivity quote  -->

## Code Review

- look at GitHub
- Call GitHub action

- headless



# Apply UI modifications between V1 and V2.

```
compare two images @current-ui.png and @updated-ui.png, find what has changed and provide the details.
```

```
locate source files in the employee-ui service and update them
```

```
in the @ViewMenuPage.vue for rating and description make the changes for values to come from the variable, like for inventory
```

[TODO - not install extension or tested it]
In fact, the site feels a bit dated, lets reamagine it 

```
nano bananna reimagine
```

go to stitch

```
I would like to expand Cymbal Eats to go into the coffee market, could you create a mockup of what the app could look like to allow me to find an awesome coffee and order it. I want to be able to track the driver bringing me my coffee
```

write a prd for this




# CodMod and Gemini CLI

```
codmod –help
```

```
codmod create -c ./cymbal-eats -o ./cymbal-eats/codmod/report.html -p coffee-and-codey -r us-central1 --modelset 2.5-pro --estimate-cost
```

```
codmod create -c ./cymbal-eats -o ./cymbal-eats/codmod/report.html -p coffee-and-codey -r us-central1 --modelset 2.5-pro --optional-sections classes
```

View the .md file

```
Create a plan for migration leveraging @codmod/report.html 
```

```
I would like to add these to jira as a new project, can you do that for me
Add each of these to a new project in jira
```

```
using @codmod/plan.json take the highest priority item and propose implementation changes
```

```
actually 1st what do I need to do to migrate from java 11 to the latest?
```

```
could you review the code and flag any early potential issues? 
```


# Playwright 

In new tab
* git clone https://github.com/GoogleCloudPlatform/testing-with-duet-ai-codelab.git
* cd testing-with-duet-ai-codelab


```
Run this app
```


```
Open the app using pupetier http://127.0.0.1:8080/ and check that text “Roman Numerals” is present and the user can enter a number and hit Convert! Button. Run several conversions(10, 25, 50) and verify results. Close the browser after you are done.
```

```
Generate and execute a script using the Puppeteer library.

The script should:
1.  Launch a browser and navigate to 'http://127.0.0.1:8080/'.
2.  Check that an element on the page contains the text "Roman Numerals".
3.  Test the following conversions:
    - Input '10' and verify the result is 'X'.
    - Input '25' and verify the result is 'XXV'.
    - Input '50' and verify the result is 'L'.
4.  Close the browser after all tests are complete.
```

```
stop the app
```

Go to the gemini app and ask to improve he prompt

```
how could this prompt be improved for gemini cli?

Generate a Node.js script using the Puppeteer library.

The script should:
1.  Launch a browser and navigate to 'http://127.0.0.1:8080/'.
2.  Check that an element on the page contains the text "Roman Numerals".
3.  Test the following conversions:
    - Input '10' and verify the result is 'X'.
    - Input '25' and verify the result is 'XXV'.
    - Input '50' and verify the result is 'L'.
4.  Close the browser after all tests are complete.
```


# Ways to use AI
* Discovery 
* Planning
* Coding
* Code Review
* Migrations
* Testing
* Bug Fixes