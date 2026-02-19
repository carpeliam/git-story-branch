# git-story-branch

*This is archived, as Pivotal Tracker is no longer with us 😭 If you're a [Shortcut](https://www.shortcut.com/) user, please consider [gitshorty](https://github.com/carpeliam/gitshorty).*

So you've got a branch, and it's in support of a particular user story, so you put your story id at the end of your branch name.

With the power of `git-story-branch`, this allows you to:

1. Open the story for the current branch in your browser
2. Delete old branches for stories that have already been accepted
3. ~~Automatically unsubscribe from unwanted spam emails~~
4. ~~Be reminded if you haven't watered your plants recently~~

At least, it will, some day.

## Installation

```sh
go build ./cli/main.go
```

## Running locally
```sh
export TRACKER_API_TOKEN="YOUR PIVOTAL TRACKER API TOKEN" # https://www.pivotaltracker.com/help/articles/api_token/
./main # assuming the library has been built and you're in a branch that has a story ID at the end
```
