import './style.css';
import './app.css';

import { Greet, SystemInfo } from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
  <main class="shell">
    <section class="hero">
      <p class="eyebrow">Go + Wails + HTML/CSS/JS</p>
      <h1>Minimal desktop app for macOS, Windows, and Linux</h1>
      <p class="lead">
        This screen is rendered with vanilla frontend code, while the button below
        calls real Go methods through Wails.
      </p>
    </section>

    <section class="panel">
      <label class="label" for="name">Your name</label>
      <div class="input-row">
        <input class="input" id="name" type="text" autocomplete="off" placeholder="Type a name" />
        <button class="btn" id="greet-btn" type="button">Ping Go</button>
      </div>
      <p class="hint">Try entering any name, then check the response coming back from Go.</p>
    </section>

    <section class="status-grid">
      <article class="card">
        <span class="card-label">Backend response</span>
        <p class="card-value" id="result">Waiting for your first call...</p>
      </article>
      <article class="card">
        <span class="card-label">Runtime info</span>
        <p class="card-value" id="system-info">Loading runtime information...</p>
      </article>
    </section>
  </main>
`;

let nameElement = document.getElementById('name');
nameElement.focus();
let resultElement = document.getElementById('result');
let systemInfoElement = document.getElementById('system-info');
let greetButton = document.getElementById('greet-btn');

async function loadSystemInfo() {
  try {
    systemInfoElement.innerText = await SystemInfo();
  } catch (err) {
    systemInfoElement.innerText = 'Could not load system info.';
    console.error(err);
  }
}

async function greet() {
  let name = nameElement.value;

  greetButton.disabled = true;
  resultElement.innerText = 'Calling Go...';

  try {
    resultElement.innerText = await Greet(name);
  } catch (err) {
    resultElement.innerText = 'Call failed. Check the console for details.';
    console.error(err);
  } finally {
    greetButton.disabled = false;
  }
}

greetButton.addEventListener('click', greet);
nameElement.addEventListener('keydown', (event) => {
  if (event.key === 'Enter') {
    greet();
  }
});

loadSystemInfo();
