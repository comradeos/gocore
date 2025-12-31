const e = React.createElement;

const DIGITS = ["0","1","2","3","4","5","6","7","8","9"];

const SYMBOLS = [".", ",", "+", "-", "*", "/"];

const KEYBOARD_LAYOUT = {
  en: {
    dir: "ltr",
    letters: [
      "q","w","e","r","t","y","u","i","o","p",
      "a","s","d","f","g","h","j","k","l",
      "z","x","c","v","b","n","m"
    ]
  },

  ru: {
    dir: "ltr",
    letters: [
      "й","ц","у","к","е","н","г","ш","щ","з","х","ъ",
      "ф","ы","в","а","п","р","о","л","д","ж","э",
      "я","ч","с","м","и","т","ь","б","ю"
    ]
  },

  he: {
  dir: "rtl",
  letters: [
    "ק","ר","א","ט","ו","ן","ם","פ",
    "ש","ד","ג","כ","ע","י","ח","ל","ך","ף",
    "ז","ס","ב","ה","נ","מ","צ","ת","ץ"
  ]
}
  
};

/* ---------- Keyboard Component ---------- */

function Keyboard({ visible, value, onChange, onOk }) {
  const [shift, setShift] = React.useState(false);
  const [lang, setLang] = React.useState("en");

  if (!visible) return null;

  const layout = KEYBOARD_LAYOUT[lang];
  const isRTL = layout.dir === "rtl";

  function addChar(c) {
    onChange(value + c);
  }

  function backspace() {
    onChange(value.slice(0, -1));
  }

  function renderKey(char, index) {
    const label = shift ? char.toUpperCase() : char;

    return e(
      "div",
      {
        key: index,
        className: "key",
        onClick: () => addChar(label)
      },
      label
    );
  }

  return e(
    "div",
    { className: "keyboard-overlay" },

    /* Input */
    e("input", {
      className: "keyboard-input",
      value: value,
      readOnly: true,
      dir: layout.dir
    }),

    /* Keys */
    e(
      "div",
      {
        className: "keyboard-keys" + (isRTL ? " rtl" : "")
      },

      /* Digits */
      DIGITS.map(renderKey),

      /* Letters */
      layout.letters.map(renderKey),

      /* Symbols */
      SYMBOLS.map(renderKey),

      /* Actions */
      e(
        "div",
        {
          className: "key key-wide key-action",
          onClick: () => setShift(!shift)
        },
        "⇧"
      ),

      e(
        "div",
        {
          className: "key key-wide key-action",
          onClick: backspace
        },
        "⌫"
      ),

      e(
        "div",
        {
          className: "key key-wide key-action",
          onClick: () => {
            const langs = Object.keys(KEYBOARD_LAYOUT);
            const idx = langs.indexOf(lang);
            setLang(langs[(idx + 1) % langs.length]);
          }
        },
        lang.toUpperCase()
      ),

      e(
        "div",
        {
          className: "key key-wide key-action",
          onClick: onOk
        },
        "OK"
      )
    )
  );
}

/* ---------- Menu ---------- */

function FullMenu({ open, onSelect }) {
  return e(
    "div",
    { className: "menu" + (open ? " open" : "") },

    e("div", { className: "menu-section" }, "Service"),
    e("div", { className: "menu-item", onClick: () => onSelect("login") }, "Login"),
    e("div", { className: "menu-item", onClick: () => onSelect("caution") }, "Caution"),
    e("div", { className: "menu-item", onClick: () => onSelect("measure") }, "Measure"),

    e("div", { className: "menu-section" }, "Settings"),
    e("div", { className: "menu-item", onClick: () => onSelect("date") }, "Date & Time"),
    e("div", { className: "menu-item", onClick: () => onSelect("network") }, "Network"),
    e("div", { className: "menu-item", onClick: () => onSelect("display") }, "Display"),
    e("div", { className: "menu-item", onClick: () => onSelect("test") }, "Test"),
  );
}

/* ---------- App ---------- */

function App() {
  const [menuOpen, setMenuOpen] = React.useState(false);
  const [page, setPage] = React.useState("home");

  const [keyboardVisible, setKeyboardVisible] = React.useState(false);
  const [keyboardValue, setKeyboardValue] = React.useState("");
  const [activeInput, setActiveInput] = React.useState(null);
  const [formData, setFormData] = React.useState({});

  function openKeyboard(inputId) {
    setActiveInput(inputId);
    setKeyboardValue(formData[inputId] || "");
    setKeyboardVisible(true);
  }

  function closeKeyboard() {
    setFormData({
      ...formData,
      [activeInput]: keyboardValue
    });
    setKeyboardVisible(false);
  }

  function renderPage() {
    switch (page) {
      case "login":
        return e(PageLogin, { openKeyboard, formData });
      case "measure":
        return e(PageMeasure);
      case "settings":
        return e(PageSettings);
      case "test":
        return e(PageTest);
      default:
        return e(PageHome);
    }
  }

  return e(
    "div",
    null,

    /* Header */
    e(
      "div",
      { className: "header" },

      e(
        "div",
        {
          className: "burger",
          onClick: () => setMenuOpen(!menuOpen)
        },
        menuOpen ? "✕" : "☰"
      ),

      e("div", { className: "header-title" }, "Caution"),

      e(
        "div",
        { className: "header-icons" },
        e("span", null, "⚡"),
        e("span", null, "🔋")
      )
    ),

    /* Menu */
    e(FullMenu, {
      open: menuOpen,
      onSelect: (p) => {
        setPage(p);
        setMenuOpen(false);
      }
    }),

    /* Content */
    !menuOpen &&
      e(
        "div",
        { className: "content" },
        renderPage()
      ),

    /* Keyboard Overlay */
    e(Keyboard, {
      visible: keyboardVisible,
      value: keyboardValue,
      onChange: setKeyboardValue,
      onOk: closeKeyboard
    })
  );
}

/* ---------- Render ---------- */

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(e(App));
















/* ---------- Pages ---------- */

function PageHome() {
  return e("div", null, "Home page");
}

function PageMeasure() {
  return e("div", null, "Measure page");
}

function PageSettings() {
  return e("div", null, "Settings page");
}

function PageLogin({ openKeyboard, formData }) {
  return e(
    "div",
    { className: "form" },

    e(
      "div",
      { className: "form-group" },
      e("label", { className: "form-label" }, "Username"),
      e("input", {
        className: "form-input",
        readOnly: true,
        value: formData.username || "",
        onClick: () => openKeyboard("username")
      })
    ),

    e(
      "div",
      { className: "form-group" },
      e("label", { className: "form-label" }, "Password"),
      e("input", {
        className: "form-input",
        readOnly: true,
        value: formData.password ? "•".repeat(formData.password.length) : "",
        onClick: () => openKeyboard("password")
      })
    )
  );
}






function Text(text) {
  return e("span", null, text);
}

function Label(text) {
  return e("span", { className: "label" }, text);
}

function Value(text) {
  return e("span", { className: "value" }, String(text));
}

function Row(label, value) {
  return e(
    "div",
    { className: "row" },
    Label(label),
    Value(value)
  );
}

function Card(title, rows) {
  return e(
    "div",
    { className: "card" },
    e("h3", null, title),
    ...rows
  );
}

function Loading() {
  return e("div", null, "Loading...");
}

function ErrorMessage(text) {
  return e("div", { className: "error" }, "Error: " + text);
}

function PageTest() {
  const [data, setData] = React.useState(null);
  const [error, setError] = React.useState(null);

  React.useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/todos/1")
      .then((r) => r.json())
      .then(setData)
      .catch((e) => setError(e.message));
  }, []);

  if (error) {
    return ErrorMessage(error);
  }

  if (!data) {
    return Loading();
  }

  return e(
    "div",
    { className: "test-page" },

    [
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
      Card("Todo item", [
        Row("User ID", data.userId),
        Row("ID", data.id),
        Row("Title", data.title),
        Row("Completed", data.completed ? "Yes" : "No")
      ]),
    ]
  );
}
