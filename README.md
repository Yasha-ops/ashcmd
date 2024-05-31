# ashcmd

`ashcmd` is a Go module designed to interact with Avast Antivirus's command-line tool `ashCmd.exe`, allowing you to run file scans and process the results in your Go applications.

## Installation

To install the module, use `go get`:

```sh
go get github.com/yasha-ops/ashcmd
```

## Usage

### Importing

First, import the module into your Go project:

```go
import (
    "github.com/yasha-ops/ashcmd"
)
```

### Initialization

Create a new instance of the `Avast` object by providing the path to `ashCmd.exe`:

```go
avast := ashcmd.NewAvast("")
// or specify the path
avast := ashcmd.NewAvast("C:\\Program Files\\Avast Software\\Avast\\ashCmd.exe")
```

If you do not provide a specific path, the default path will be used.

### Running a Scan

To run a file scan with `ashCmd.exe` and get the results, use the `ScanFile` method:

```go
result, err := avast.ScanFile()
if err != nil {
    fmt.Println("Error executing the scan:", err)
    return
}

if result.IsDetected {
    fmt.Println("Malware detected:", result.Malware)
    fmt.Println("Engine version:", result.EngineVersion)
} else {
    fmt.Println("No malware detected.")
}
```

## Result Structure

The `ScanResult` structure contains the following information:

- `IsDetected` (bool): Indicates whether malware was detected.
- `Malware` (string): The name of the detected malware (if applicable).
- `EngineVersion` (string): The version of the Avast detection engine used during the scan.

## Complete Example

Here is a complete example of using the `ashcmd` module:

```go
package main

import (
    "fmt"
    "github.com/yasha-ops/ashcmd"
)

func main() {
    avast := ashcmd.NewAvast("")

    result, err := avast.ScanFile()
    if err != nil {
        fmt.Println("Error executing the scan:", err)
        return
    }

    if result.IsDetected {
        fmt.Println("Malware detected:", result.Malware)
        fmt.Println("Engine version:", result.EngineVersion)
    } else {
        fmt.Println("No malware detected.")
    }
}
```

## Contributions

Contributions are welcome! If you have any improvement ideas or find any bugs, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.