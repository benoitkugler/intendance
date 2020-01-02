import { Controller } from "./controller";
import Cookie from "js-cookie";
import { OutLoggin, InLoggin } from "./types";
import axios, { AxiosResponse } from "axios";
import { ServerURL } from "./data";

export class Loggin {
  private readonly controller: Controller;

  constructor(c: Controller) {
    this.controller = c;
  }

  private checkCookies() {
    const token = Cookie.get("token");
    const idUtilisateur = Cookie.get("id_utilisateur");
    if (token == undefined || idUtilisateur == undefined) {
      return null;
    }
    return { token, idUtilisateur };
  }

  // renvoie un message d'erreur ou la chaine vide
  // si le mot de passe est correct.
  async loggin(params: InLoggin) {
    try {
      const response: AxiosResponse<OutLoggin> = await axios.post(
        ServerURL + "/loggin",
        params
      );
      if (response.data.erreur != "") {
        return response.data.erreur;
      }
      this.controller.token = response.data.token;
      this.controller.idUtilisateur = response.data.utilisateur.id;
      this.controller.notifications.setMessage(
        `Connecté sous le nom de <b>${response.data.utilisateur.prenom_nom}</b>`
      );
      this.controller.isLoggedIn = true;
    } catch (error) {
      return this.controller.notifications.setAxiosError(error);
    }
  }

  logout() {
    this.controller.isLoggedIn = false;
    // TODO: cookies
  }
}
