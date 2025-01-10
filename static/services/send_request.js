export async function sendRequest(url, method, errorMessage, body) {
  let bodyResponse
  let errorResponse

  try {
    const response = await fetch(url, {
      method: method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body)
    });

    if (!response.ok) throw new Error(errorMessage);

    const contentType = response.headers.get('Content-Type');

    if (contentType?.includes('application/json')) {
      bodyResponse = await response.json();
    } else {
      bodyResponse = await response.text();
    }
  } catch (error) {
    console.error(error);
    errorResponse = error
  }

  return new RequestResponse(bodyResponse, errorResponse);
}

export class RequestResponse{
  constructor(body, error) {
    this.body = body
    this.error = error
  }

  hasError() {
    return !!this.error;
  }

  getErrorMessage() {
    return this.error.message;
  }

  getBody() {
    return this.body;
  }
}