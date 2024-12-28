async function saveBankAccount(event) {
  event.preventDefault();

  try {
    const valueToSend = {
      acronym_value: document.getElementById('acronymValue').value.trim(),
      description_value: document.getElementById('descriptionValue').value.trim(),
    };

    if (!valueToSend.acronym_value) {
      throw new Error('A sigla é obrigatória!');
    }

    const response = await fetch('/api/save-account', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(valueToSend),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Erro ao salvar os dados. Por favor, tente novamente.');
    }


    document.getElementById('acronymSave').innerText = valueToSend.acronym_value;
    document.getElementById('descriptionSave').innerText = valueToSend.description_value;


    document.getElementById('accountForm').reset();

    document.getElementById('errorMessage').innerText = '';
  } catch (error) {
    console.error(error);
    document.getElementById('errorMessage').innerText = error.message;
  }
}

async function getAccounts() {
  try {
    const response = await fetch('/api/get-accounts', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {
      throw new Error(`Error fetching accounts: ${response.status}`);
    }

    const data = await response.json();

    const tbody = document.querySelector("table tbody");
    tbody.innerHTML = "";

    data.forEach((account) => {
      const row = document.createElement("tr");

      const bankCell = document.createElement("td");
      bankCell.textContent = account.bank;
      row.appendChild(bankCell);

      const accountCell = document.createElement("td");
      accountCell.textContent = account.account;
      row.appendChild(accountCell);

      const acronymCell = document.createElement("td");
      acronymCell.textContent = account.acronym;
      row.appendChild(acronymCell);

      tbody.appendChild(row);
    });
  } catch (error) {
    console.error(error);
    document.getElementById('errorMessage').innerText = error.message;
  }
}

document.addEventListener("DOMContentLoaded", getAccounts);