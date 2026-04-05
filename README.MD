# 🎨 ASCII Art Generator (Go Web App)

A web-based ASCII art generator built with **Go (Golang)** that converts user input text into styled ASCII art with support for:

* Multiple banner fonts
* Substring highlighting
* Custom colors
* File download functionality

---

## 🚀 Features

* 📝 Convert plain text to ASCII art
* 🎭 Choose from multiple banners:

  * `standard`
  * `shadow`
  * `thinkertoy`
* 🎨 Highlight specific substrings with custom colors
* 📥 Download generated ASCII art as a `.txt` file
* 🌐 Simple and interactive web interface
* ⚡ Fast server-side rendering using Go

---

## 🏗️ Project Structure

```
.
├── main.go                  # Entry point (HTTP server & handlers)
├── go.mod                  # Go module file
├── ascii/
│   ├── generator.go        # ASCII generation + coloring logic
│   └── banners/            # Banner font files
│       ├── standard.txt
│       ├── shadow.txt
│       └── thinkertoy.txt
├── templates/
│   ├── index.html          # Main UI
│   └── error.html          # Error page
├── static/
│   └── style.css           # Styling
```

---

## ⚙️ Installation & Setup

### 1. Clone the repository

```
git clone https://github.com/your-username/ascii-art-generator.git
cd ascii-art-generator
```

---

### 2. Install dependencies

```
go mod tidy
```

---

### 3. Run the server

```
go run .
```

---

### 4. Open in browser

```
http://localhost:8001
```

---

## 🧠 How It Works

### 1. Input Handling

* User submits text, banner type, substring, and color via a form.

### 2. ASCII Generation

* Each character is mapped to a predefined ASCII pattern from banner files.
* The app reads banner files line-by-line and constructs output row-by-row.

### 3. Substring Highlighting

* A mask is built to detect substring positions.
* Matching characters are wrapped in HTML `<span>` tags with inline color styles.

### 4. Rendering

* Output is rendered inside a `<pre>` tag using safe HTML injection.

---

## 🎯 Example Usage

Input:

```
Text: Hello
Substring: ll
Color: red
Banner: standard
```

Output:

* ASCII art of "Hello"
* Letters "ll" highlighted in red

---

## 📥 Download Feature

* Click **Download Art**
* The generated ASCII is returned as a `.txt` file via HTTP response headers

---

## ⚠️ Error Handling

* Invalid request methods return appropriate HTTP status codes
* Empty input is validated
* File read errors are handled gracefully

---

## 🛠️ Technologies Used

* Go (net/http, html/template)
* HTML5 & CSS3
* ASCII banner files

---

## 📌 Future Improvements

* 🎨 Color picker UI instead of manual input
* 💾 Save generated art history
* 🌈 Gradient or multi-color support
* ⚡ Caching banner files for performance

---

## 👨‍💻 Author

Built by **Bolarinwa Oladimeji Emmanuel**

---

## 📄 License

This project is open-source and available under the MIT License.
