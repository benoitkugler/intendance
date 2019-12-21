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

class NotificationsState {
  private error: Error | null;
  private message: string | null;
  private spin: boolean;

  constructor() {
    this.error = null;
    this.message = null;
    this.spin = false;
  }

  getError() {
    return this.error;
  }

  getMessage() {
    return this.message;
  }

  getSpin() {
    return this.spin;
  }

  startSpin() {
    this.spin = true;
  }

  setMessage(message: string | null) {
    this.spin = false;
    this.message = message;
  }

  setError(error: Error | null) {
    this.spin = false;
    this.error = error;
  }

  setAxiosError(error: any) {
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
          messageHtml = `Impossible de décoder la réponse du serveur. <br/>
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
    this.error = { kind, messageHtml, code };
    this.spin = false;
  }
}

// Contient les données nécessaires aux notifications
// (erreurs ou succès)
export const NS = new NotificationsState();
