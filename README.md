# TODO Search
This is a simple script which when executed in a directory, produces a list of all files (using their absolute paths) containing the keyword "TODO" in them. The files can be in the immediate directory, or a sub-directory (or a sub-directory of the sub-directory, ad infinitum).

## Built with
Go 🚀

## Getting Started

### Prerequisites
Ensure that you have Go installed. Otherwise, click [here](https://golang.org/doc/install) to download and install Go.

### Clone the Project

```bash
$ git clone https://github.com/yuqiil/mcf-todo
```

### Run the script
```bash
$ cd mcf-todo
$ go run main.go
```

## Features
- Developed for searching javascript files, however it can also be modified to search for any file or other file type. 
- Developed based on the assumption that it will only flag out the exact string ```"TODO"``` in upper case. This is done through regular expression.
- This script can also be modified to search for other keywords.

## Usage
Given the following directory structure:
```bash
/path/to/your/dir
  - somedir
    - somemodule
      - somefile.js
      - someotherfile.js
  - somedir2
    - anotherdir
      - yet_another_dir
        - index.js
      - index.js
    - index.js
  - somedir3
    - another_file.js
```
Assuming the application runs at ```/path/to/your/dir```, and assuming all of the files contains the text string ```"TODO"```, the application will output something similar to:
```
/path/to/your/dir/somedir/somemodule/somefile.js
/path/to/your/dir/somedir/somemodule/someotherfile.js
/path/to/your/dir/somedir2/anotherdir/yet_another_dir/index.js
/path/to/your/dir/somedir2/anotherdir/index.js
/path/to/your/dir/somedir2/index.js
/path/to/your/dir/somedir3/another_file.js
```

## License
Content is licensed under the [Creative Commons 4.0 (Non-Commercial, Attribution) license](https://creativecommons.org/licenses/by-nc-sa/4.0/)
