import { RepasWithGroupe } from "@/logic/types";

export interface DateTime {
  date: string;
  time: string;
}

export interface DataEvent {
  start: Time;
  repas: RepasWithGroupe;
}

export function toDateVuetify(d: Date) {
  return d.toISOString().substr(0, 10);
}

export function formatNbOffset(repas: RepasWithGroupe) {
  const n = repas.offset_personnes;
  if (n != 0) {
    return `${n > 0 ? "+" : ""}${n} pers.`;
  }
  return "";
}

// function findClosest<K extends { distance: number }>(values: K[]) {
//   let bestV: K | null = null;
//   let bestDist: number = Infinity;
//   for (const v of values) {
//     if (v.distance < bestDist) {
//       bestV = v;
//       bestDist = v.distance;
//     }
//   }
//   return bestV;
// }

// function mtoH(minutes: number): Horaire {
//   const m = minutes % 60;
//   return { heure: (minutes - m) / 60, minute: m };
// }
// const hToM = (h: Horaire) => h.heure * 60 + h.minute;
// const _centers = Horaires.map(v => hToM(v.value)).sort((a, b) =>
//   a > b ? 1 : -1
// );

// export function repartitRepas(
//   repas: RepasWithGroupe[],
//   startTime: Horaire,
//   dureeJour: number,
//   dureeRepasMax: number
// ) {
//   // on regroupe les repas par zone
//   const groupes: { [key: number]: RepasWithGroupe[] } = {};
//   repas.forEach(r => {
//     const tmp = _centers.map(h => {
//       return {
//         horaire: h,
//         distance: Math.abs(h - hToM(r.horaire))
//       };
//     });
//     const best = findClosest(tmp)!;
//     let l = groupes[best.horaire] || [];
//     l.push(r);
//     groupes[best.horaire] = l;
//   });

//   const N = repas.length;
//   const K = _centers.length;
//   // on ne dépasse pas la hauteur max
//   const repasDuree = Math.min(dureeRepasMax, dureeJour / N);
//   // on calcule l'écart entre chaque groupe
//   const ecartGroupes = (dureeJour - repasDuree * N) / (K > 1 ? K - 1 : 1);
//   // on attribue l'horaire final

//   const out: RepasWithGroupe[] = [];
//   let currentPos = hToM(startTime);
//   for (let i = 0; i < K; i++) {
//     const horaireCentre = _centers[i];
//     const repasGroupe = (groupes[horaireCentre] || []).sort((a, b) =>
//       hToM(a.horaire) < hToM(b.horaire) ? 1 : -1
//     );
//     for (let rep of repasGroupe) {
//       console.log(currentPos);
//       rep = JSON.parse(JSON.stringify(rep));
//       rep.horaire = mtoH(currentPos);
//       out.push(rep);
//       currentPos += repasDuree;
//     }
//     if (i < K - 1) {
//       const bestNext = _centers[i + 1];
//       if (currentPos + ecartGroupes > bestNext) {
//         // on remonte le prochain groupe, car il y a la place
//         currentPos = bestNext;
//       } else {
//         currentPos += ecartGroupes;
//       }
//     } else {
//       currentPos += ecartGroupes;
//     }
//   }
//   return out;
// }
