import { sendRequest } from '../../services/send_request.js';

    if (!valueToSend.acronymValue) {
      throw new Error('A sigla é obrigatória!');
    }

    const response = await fetch('/api/save-bank-account', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(valueToSend),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Erro ao salvar os dados. Por favor, tente novamente.');
    }

    await getBankAccounts();

    document.getElementById('accountForm').reset();
    document.getElementById('errorMessage').innerText = '';
  } catch (error) {
    console.error(error);
    document.getElementById('errorMessage').innerText = error.message;
  }
}

async function getBankAccounts() {
  const response = await sendRequest('/api/get-bank-accounts', 'GET');

  if(response.hasError()) setErrorMessage(response.getErrorMessage());
  else createBankAccountTableRows(response.getBody());
}

function setErrorMessage(message) {
  document.getElementById('errorMessage').innerText = message;
}

function createBankAccountTableRows(data) {
  const tbody = document.getElementById("bankAccountRow");
  tbody.innerHTML = "";

  data.forEach((account) => {
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

document.addEventListener("DOMContentLoaded", getBankAccounts);