export interface Error {
  code: number | null;
  kind: string;
  messageHtml: string;
}

function arrayBufferToString(buffer: ArrayBuffer) {
  const uintArray = new Uint8Array(buffer);
  const encodedString = String.fromCharCode.apply(null, Array.from(uintArray));
  return decodeURIComponent(escape(encodedString));
}

// Contient les données nécessaires aux notifications
// (erreurs ou succès)
export class Notifications {
  private error: Error | null = null;
  private _messages: { text: string; id: number }[] = [];
  private spin: boolean = false;

  private queueIndex = 0;

  getError() {
    return this.error;
  }

  get messages() {
    return this._messages.map(v => v.text);
  }

  getSpin() {
    return this.spin;
  }

  startSpin() {
    // on enlève une éventuelle notication
    this.spin = true;
  }

  setMessage(message: string) {
    this.spin = false;
    const index = this.queueIndex;
    this.queueIndex++;
    this._messages.push({ text: message, id: index });
    setTimeout(() => {
      this._messages = this._messages.filter(v => v.id != index);
    }, 4000);
  }
  clearMessages() {
    this._messages = [];
  }

  setError(error: Error | null) {
    this.spin = false;
    this.error = error;
  }

  private parseAxiosError(error: any): Error {
    let kind: string,
      messageHtml: string,
      code = null;
    if (error.response) {
      // The request was made and the server responded with a status code
      // that falls out of the range of 2xx
      kind = `Erreur côté serveur`;
      code = error.response.status;

      messageHtml = error.response.data.message;
      if (!messageHtml) {
        try {
          const json = arrayBufferToString(error.response.data);
          messageHtml = JSON.parse(json).message;
        } catch (error) {
          messageHtml = `Le format d'erreur du serveur n'a pu être décodé.<br/>
        Détails : <i>${error}</i>`;
        }
      }
    } else if (error.request) {
      // The request was made but no response was received
      // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
      // http.ClientRequest in node.js
      kind = "Aucune réponse du serveur";
      messageHtml =
        "La requête a bien été envoyée, mais le serveur n'a donné aucune réponse...";
    } else {
      // Something happened in setting up the request that triggered an Error
      kind = "Erreur du client";
      messageHtml = `La requête n'a pu être mise en place. <br/>
                  Détails :  ${error.message} `;
    }
    return { kind, messageHtml, code };
  }

  setAxiosError(error: any) {
    this.error = this.parseAxiosError(error);
    this.spin = false;
  }
}
