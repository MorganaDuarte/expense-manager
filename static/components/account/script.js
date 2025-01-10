import {RequestResponse } from '../../services/RequestResponse.js';

async function getBankAccounts() {
  const request = new RequestResponse();
  const response = await request.sendRequest('/api/get-bank-accounts', 'GET');

  if(response.hasError()) setErrorMessage(response.getErrorMessage());
  else createBankAccountTableRows(response.getBody());
}

async function saveBankAccount(event) {
  event.preventDefault();

  const body = {
    acronym: document.getElementById('acronymValue').value.trim(),
    description: document.getElementById('descriptionValue').value.trim(),
  };

  if (!body.acronym) throw new Error('A sigla é obrigatória!');

  const request = new RequestResponse();
  const response = await request.sendRequest('/api/save-bank-account', 'POST', "", body);
  if(response.hasError()) {
    setErrorMessage(response.getErrorMessage());
  } else {
    await getBankAccounts();
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