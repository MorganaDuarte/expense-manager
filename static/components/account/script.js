import RequestSender from "../../services/RequestSender.js";

let bankAccounts = [];

async function getBankAccounts() {
  const requestSender = new RequestSender();
  const response = await requestSender.send('/api/get-bank-accounts', 'GET');

  if(response.hasError()) {
    setErrorMessage(response.getErrorMessage());
  } else {
    bankAccounts = response.getBody();
    createBankAccountTableRows(bankAccounts);
  }
}

async function saveBankAccount(event) {
  event.preventDefault();

  const body = {
    acronym: document.getElementById('acronymValue').value.trim(),
    description: document.getElementById('descriptionValue').value.trim(),
  };

  if (!body.acronym) throw new Error('A sigla é obrigatória!');

  const requestSender = new RequestSender();
  const response = await requestSender.send('/api/save-bank-account', 'POST', "", body);
  if(response.hasError()) {
    setErrorMessage(response.getErrorMessage());
  } else {
    bankAccounts.push({acronym: body.acronym, description: body.description});
    createBankAccountTableRows(bankAccounts);
    document.getElementById('accountForm').reset();
  }
}

function setErrorMessage(message) {
  document.getElementById('errorMessage').innerText = message;
}

function createBankAccountTableRows(data) {
  const tbody = document.getElementById("bankAccountRow");
  tbody.innerHTML = "";

  data?.forEach((account) => {
    const row = document.createElement("tr");

    const acronymCell = document.createElement("td");
    acronymCell.textContent = account.acronym;
    row.appendChild(acronymCell);

    const descriptionCell = document.createElement("td");
    descriptionCell.textContent = account.description;
    row.appendChild(descriptionCell);

    tbody.appendChild(row);
  });
}

document.addEventListener("DOMContentLoaded", () => {
  getBankAccounts();
  document.getElementById('accountForm').addEventListener('submit', saveBankAccount);
});