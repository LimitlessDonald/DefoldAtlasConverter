# DefoldAtlasConverter
DefoldAtlasConverter is a simple command line application to convert the atlas file from the [defold](https://defold.com) game engine into various formats

## Features
* Convert .atlas file to .json file
* Export/Convert the images/image names in a .atlas file into lua tables 

## Installation & Usage

This tool should work on most operating systems including Windows, macOS and Linux.


### Simple Installation
I would upload executables soon, so you don't have to compile anything. But compiling from source is super easy !

### Semi-Advanced Installation (From source)
Dependencies : You need to have golang installed . It's very easy to install https://go.dev/doc/install

#### Linux & macOS :

```shell
 git clone https://github.com/LimitlessDonald/DefoldAtlasConverter
```
```shell
 cd DefoldAtlasConverter
```

```shell
  ./install.sh
```
***NOTE:*** If for some reasons the installation script doesn't work for you, the method below should work just fine 
#### Windows : 
```shell
 git clone https://github.com/LimitlessDonald/DefoldAtlasConverter
```

```shell
 cd DefoldAtlasConverter
```

```shell
 go build defoldAtlasConv.go
```

On Windows a `defoldAtlasConv.exe` file should be generated . Its just `"defoldAtlasConv"` for Linux and macOS.

To use the application in Windows command line from any location, you would need to add the folder where the application exists to your PATH 


### Usage

#### Linux & macOS :

To export to JSON
```shell
defoldAtlasConv -atlas test.atlas -output test_output.json
```
OR 
```shell
defoldAtlasConv -atlas=test.atlas -output=test_output.json
```

To export to Lua file, containing a table of the image names
```shell
defoldAtlasConv -atlas test.atlas -output test_output.lua
```
OR
```shell
defoldAtlasConv -atlas=test.atlas -output=test_output.lua
```

**NOTE:** If you don't provide an output file, it would automatically write to `output.lua`


##Support, Follow, or Hire Me
You can support me and my work by contributing to this repo, donating if you can or following me on any of my social platforms via the links below 

If you would also like to hire me, you can send me a DM on any of these platforms

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png)](https://www.buymeacoffee.com/limitlessDonald)

[My Website](https://limitlessdonald.com) (You can learn more about me on my personal website)

[Twitter](https://twitter.com/LimitlessDonald)
[Facebook](https://facebook.com/LimitlessDonald)
[Instagram](https://instagram.com/LimitlessDonald)
[Reddit](https://www.reddit.com/user/LimitlessDonald)