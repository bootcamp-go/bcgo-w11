## Understanding `os.OpenFile` in Go

The `os.OpenFile` function in Go is a versatile way to open a file with specific flags and permissions. Its signature is:

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

### Parameters

- **`name string`**: The file path to open or create.
- **`flag int`**: An integer representing the file mode flags (e.g., read-only, write-only, read-write, append, create, truncate).
> flags include:
> - os.O_RDONLY: Read-only mode.
> - os.O_WRONLY: Write-only mode.
> - os.O_RDWR: Read-write mode.
> - os.O_APPEND: Append to the file when writing.
> - os.O_CREATE: Create a new file if it doesn't exist.
> - os.O_TRUNC: Truncate the file when opening.
> These flags can be combined using bitwise OR (|). For example, os.O_CREATE | os.O_WRONLY opens the file for write-only access, creating the file if it does not exist.
- **`perm FileMode`**: The file permission code.

### File Permission Codes

The `perm FileMode` is a Unix-style file permission, typically specified in octal format. It's composed of 4 digits, where each digit represents different sets of permissions:

1. **Special Permissions**: The first digit is used for special permissions and is often set to 0 in common use cases.
2. **Owner Permissions**: The second digit represents permissions for the file owner.
3. **Group Permissions**: The third digit represents permissions for the group.
4. **Others Permissions**: The fourth digit represents permissions for others.

### Permission Values

- `0`: No permissions.
- `1`: Execute only.
- `2`: Write only.
- `4`: Read-only.

### Permission Combinations
- `3`: Write and execute.
- `5`: Read and execute.
- `6`: Read and write.
- `7`: Read, write, and execute.

### Most Common Permission Codes
- `0644`: Owner can read and write. Group and others can read.
- `0755`: Owner can read, write, and execute. Group and others can read and execute.

### Example

If we use `0644` as the permission code in `os.OpenFile`:

```go
os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY, 0644)
```

This means:

- The file "example.txt" is opened with write-only access, and created if it does not exist.
- The owner of the file has read and write permissions (`6`).
- The group and others have read-only permissions (`4` for both).

Remember, these permissions apply only when the file is being created. If the file exists, `OpenFile` will not change its existing permissions.

To change the permissions of an existing file, you would use the os.Chmod function. The os.Chmod function allows you to change the file permissions at any time, not just at the time of file creation. Its usage is as follows:

```go
err := os.Chmod("example.txt", 0644)
if err != nil {
    // handle error
}
```

This will change the permissions of the file named "filename" to read-write for the owner and read-only for the group and others, regardless of the file's existing permissions. Remember that the permission code in os.Chmod is also specified in octal format, similar to os.OpenFile.

In the file system it would look like this

```bash
$ ls -l
-rw-r--r-- example.txt
```