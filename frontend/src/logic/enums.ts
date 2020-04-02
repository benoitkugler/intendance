
	export interface EnumItem<T = string> {
		value: T;
		text: string;
  	}
	
	function fmt(enums: EnumItem[], value: string) {
		const data = enums.find(v => v.value == value);
		if (data) return data.text;
		return "";
	}

	
	export const HoraireFields = {PetitDejeuner: "matin",
		Midi: "midi",
		Gouter: "gouter",
		Diner: "diner",
		Cinquieme: "cinquieme",
		};
  	export const Horaires : EnumItem[] = [{ value: HoraireFields.PetitDejeuner, text: "Petit déjeuner" },
		  { value: HoraireFields.Midi, text: "Midi" },
		  { value: HoraireFields.Gouter, text: "Goûter" },
		  { value: HoraireFields.Diner, text: "Dîner" },
		  { value: HoraireFields.Cinquieme, text: "Cinquième" },
		  ];
	export const fmtHoraire = (v: string) => fmt(Horaires, v);

	
	export const UniteFields = {Litres: "L",
		Kilos: "Kg",
		Piece: "P",
		};
  	export const Unites : EnumItem[] = [{ value: UniteFields.Litres, text: "Litres" },
		  { value: UniteFields.Kilos, text: "Kilos" },
		  { value: UniteFields.Piece, text: "Pièce(s)" },
		  ];
	export const fmtUnite = (v: string) => fmt(Unites, v);

	
