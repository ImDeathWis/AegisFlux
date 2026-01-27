const toggle = document.getElementById('mode-toggle');
const body = document.body;

// Toggle modo oscuro / claro
toggle.addEventListener('click', () => {
  if(body.classList.contains('dark-mode')){
    body.classList.replace('dark-mode', 'light-mode');
    toggle.textContent = '☀️';
  } else {
    body.classList.replace('light-mode', 'dark-mode');
    toggle.textContent = '🌙';
  }
});

// Simular datos de dashboard
const dashboardData = [
  {name: "Agentes activos", value: 12},
  {name: "Conexiones recientes", value: 54},
  {name: "Alertas críticas", value: 2, alert: true}
];

const dashboardContainer = document.getElementById('dashboard-cards');
dashboardData.forEach(item => {
  const card = document.createElement('div');
  card.className = 'card' + (item.alert ? ' alert' : '');
  card.innerHTML = `${item.name}<br><span class="value">${item.value}</span>`;
  dashboardContainer.appendChild(card);
});

// Simular logs
const logsData = [
  {time: "12:01", agent: "Agent1", msg: "Error: conexión perdida"},
  {time: "12:02", agent: "Agent2", msg: "Info: todo normal"},
  {time: "12:05", agent: "Agent3", msg: "Alerta: actividad sospechosa"},
];

const logsTable = document.getElementById('logs-table');
logsData.forEach(log => {
  const tr = document.createElement('tr');
  tr.innerHTML = `<td>${log.time}</td><td>${log.agent}</td><td>${log.msg}</td>`;
  logsTable.appendChild(tr);
});
