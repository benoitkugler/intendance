import { Controller } from "./controller";
import axios, { AxiosResponse } from "axios";
import { ServerURL } from "./data";
import { OutResoudIngredients, InResoudIngredients } from "./types";

export class Calculs {
  private controller: Controller;

  constructor(controller: Controller) {
    this.controller = controller;
  }

  async resoudIngredientsRepas(idRepas: number, nbPersonnes?: number) {
    const params: InResoudIngredients = {
      mode: "repas",
      id_repas: idRepas,
      nb_personnes: nbPersonnes == undefined ? -1 : nbPersonnes,
      id_sejour: -1,
      jour_offset: -1
    };
    try {
      const response: AxiosResponse<OutResoudIngredients> = await axios.post(
        ServerURL + "/resolution",
        params,
        {
          auth: this.controller.auth()
        }
      );
      this.controller.token = response.data.token;
      return response.data;
    } catch (error) {
      this.controller.notifications.setAxiosError(error);
    }
  }
}
