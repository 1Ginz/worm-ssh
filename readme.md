# Base on [GGH](https://github.com/byawitz/ggh)

### What is GGH?

GGH is a lightweight, fast wrapper around your SSH commands. It helps you to recall your SSH sessions.
This is one of the most useful tools for developers who work with multiple servers.

Thanks to GGH Team.

### What is Worm?

Worm inherits all the features of GGH and expands with several new capabilities (I need), such as:

- Setting up a workspace for each project or company...
- Each workspace contains multiple files, with each file managing a separate list of servers that need SSH access. A
  file can represent a DataCenter or Platform…
- Supporting server access via SSH and TSH
- Saving history for each workspace

### Installation && Configuration

- Clone the repository
- Run command `go install .`
- Setup Conflict workspace

```shell
cd ~/.worm
mkdir workspace1
cd workspace1 
vim DC1 
```

- Setup Configuration file

```text
Host {HostNAme}
	HostName {IP}
	User {UserName}
	Mode {SSH|TSH}
````

### Usage

List workspaces

```shell
worm --workspace
```

Switch workspace

```shell
worm --active 
```

Interactive history

```shell
worm 
```

```shell
worm --history
```

Interactive configuration file

```shell
worm -
```

Interactive configuration file with search for groups or hostnames.

```shell
worm - xxxx
```
