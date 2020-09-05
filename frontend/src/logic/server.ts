import Cookie from "js-cookie";
import { AbstractAPI } from "./api";
import { Notifications } from "./notifications";
import { Controller } from "./controller";
import { InLoggin } from "./types";
export const devMode = process.env.NODE_ENV != "production";
const host = devMode ? "http://localhost:1323" : window.location.origin;

export class API extends AbstractAPI {
  constructor(private notifications: Notifications, token: string) {
    super(host, token, {});
  }

  handleError(err: any) {
    this.notifications.setAxiosError(err);
  }
}

export interface Meta {
  idUtilisateur: number;
  token: string;
}

const metaDev: Meta = {
  idUtilisateur: 2,
  token:
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZFByb3ByaWV0YWlyZSI6MiwiZXhwIjoxNTk5NTgzNDU5fQ.FUnzS7wMn5nJB-bqvCOhdqnXDTwvSRBBLLX05QoaG98"
};

export class Loggin {
  constructor(private readonly notifications: Notifications) {
    // const aut = this.checkCookies();
    //TODO: handle loggin via cookie
    // if (aut != null) {
    //   this.controller.token = aut.token;
    //   this.controller.idUtilisateur = aut.idUtilisateur;
    //   this.controller.state.isLoggedIn = true;
    // }
  }

  /** Vérifie les cookies */
  checkCookies(): Meta | null {
    const token = Cookie.get("token");
    const idUtilisateur = Cookie.get("id_utilisateur");
    if (token == undefined || idUtilisateur == undefined) {
      return null;
    }
    return { idUtilisateur: Number(idUtilisateur), token };
  }

  // renvoie un message d'erreur ou la chaine vide
  // si le mot de passe est correct.
  async loggin(params: InLoggin) {
    const out = await new API(this.notifications, "").Loggin(params); // token is ignored here
    if (out === undefined) return; // erreur inatendue

    // save for future connections
    Cookie.set("token", out.token);
    Cookie.set("id_utilisateur", out.utilisateur);

    this.notifications.setMessage(
      `Connecté sous le nom de <b>${out.utilisateur.prenom_nom}</b>`
    );
    return out;
  }

  logout() {
    Cookie.remove("token");
    Cookie.remove("id_utilisateur");
  }
}

export const N = new Notifications();

// point d'entrée : le controller n'est pas encore disponible
export function init() {
  const meta = new Loggin(N).checkCookies();
  if (meta !== null) {
    return new Controller(meta, N);
  }
  if (devMode) {
    return new Controller(metaDev, N);
  }
}
