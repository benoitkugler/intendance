import { AgendaUtilisateur, OutAgenda } from "./types";
import axios, { AxiosResponse } from "axios";

const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;
export const ServerURL = host + "/api";

type Cb = () => void;

interface Error {
  code: number | null;
  kind: string;
  messageHtml: string;
}

function arrayBufferToString(buffer: ArrayBuffer) {
  const uintArray = new Uint8Array(buffer);
  const encodedString = String.fromCharCode.apply(null, Array.from(uintArray));
  return decodeURIComponent(escape(encodedString));
}

function formateError(error: any): Error {
  let kind: string,
    messageHtml: string,
    code = null;
  if (error.response) {
    // The request was made and the server responded with a status code
    // that falls out of the range of 2xx
    kind = `Erreur côté serveur`;
    code = error.response.status;
    console.log(error.response.data);

    messageHtml = error.response.data.message;
    if (!messageHtml) {
      const json = arrayBufferToString(error.response.data);
      messageHtml = JSON.parse(json).message;
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

// Object principal de stockage des données
// sur le client.
// Une instance de cet objet est créé au chargement,
// puis partagée entre les différents composants.
// Le système de réactivité de vuejs permet de propager
// facilement les changements effectués aux données.
// Ce composant est responsable de la comunication avec le serveur.
export class Data {
  agenda: AgendaUtilisateur;
  error: Error | null;
  private token: string;
  private idUtilisateur: number | "*" | null;

  constructor() {
    this.agenda = { sejours: [] };
    this.token = "";
    this.idUtilisateur = devMode ? "*" : null;
    this.error = null;
  }

  private auth() {
    return {
      username: String(this.idUtilisateur || ""),
      password: this.token
    };
  }

  async loadAgenda() {
    await axios
      .get(ServerURL + "/agenda", {
        auth: this.auth()
      })
      .then((response: AxiosResponse<OutAgenda>) => {
        this.token = response.data.token;
        this.agenda = response.data.agenda;
      })
      .catch(error => {
        this.error = formateError(error);
      });
  }
}
