# CLI Todo List (Golang)

## Description

A Todo List CLI tool.  
The purpose of this project is to develop non-web CRUD ops in Golang.

## Commands

### add \<task>

- *adds task to todo list*

***Example:***

`% tasks add "brush teeth"`

will add `brush teeth` to the database

---

### list [-a]
- *returns a list of all **uncompleted** tasks*
- *using flag `-a` returns all tasks, including completed tasks*

***Example 1***

`% tasks list`

will return

```
ID      Task        Created
1       task        a minute ago
3       task2       a few seconds ago
```

***Example 2***

`% tasks list -a`

will return

```
ID      Task        Created                 Done
--      ----        -------                 ----
1       task        2 minutes ago           false
2       task2       a minute ago            true
3       task3       a few seconds ago       false
```
---

### complete \<taskID>

- *changes task with taskID to `true` in "Done" column*

***Example***

`% tasks list -a`

```
ID      Task        Created                 Done
--      ----        -------                 ----
1       task        2 minutes ago           false
2       task2       a minute ago            true
3       task3       a few seconds ago       false
```

`% tasks complete 1`

`% tasks list -a`

```
ID      Task        Created                 Done
--      ----        -------                 ----
1       task        3 minutes ago           true
2       task2       2 minutes ago           true
3       task3       a minute ago            false
```
---
### delete \<taskID>

- *removes task with taskID from todo list*

***Example***

`% tasks list -a`

```
ID      Task        Created                 Done
--      ----        -------                 ----
1       task        2 minutes ago           false
2       task2       a minute ago            true
3       task3       a few seconds ago       false
```

`% tasks delete 1`

`% tasks list -a`

```
ID      Task        Created                 Done
--      ----        -------                 ----
2       task2       2 minutes ago           true
3       task3       a minute ago            false
```

## Packages

- **[encoding/csv]** for writing out as a csv file
- **[strconv]** for turning types into strings and visa versa
- **[text/tabwriter]** for writing out tab aligned output
- **[os]** for opening and reading files
- **[github.com/spf13/cobra]** for the command line interface
- **[github.com/mergestat/timediff]** for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)


