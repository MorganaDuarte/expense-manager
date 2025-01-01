export async function sendRequest(url, method, errorMessage, body) {
  let bodyResponse
  let errorResponse

  try {
    const response = await fetch(url, {
      method: method,
      headers: { 'Content-Type': 'application/json' },
      body: body
    });

    if (!response.ok) {
      throw new Error(errorMessage);
    }

    bodyResponse = await response.json();
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