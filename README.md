# worktree-manager

## Why
For removing Git worktrees in bulk. 
Scans files in worktree, finding last modified date to help selection process.

## How
### Build executable:
```
go build -o {What you want to build the file as}

Example: 
go build -o wm
```

### In .zshrc

```
export PATH="$PATH:{REPO DIR}"

Example:
export PATH="$PATH:/Users/example/worktree-manager"
```

### Use:

- cd to worktree root dir
- Run the command name

```
Example: 
wm
```
