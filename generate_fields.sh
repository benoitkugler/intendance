../struct2ts/main -H -i -D -o frontend/src/logic/types.ts \
    views.OutIngredient views.OutIngredients \
    views.OutRecette views.OutRecettes \
    views.OutMenu views.OutMenus \
    views.OutSejour views.OutSejours \
    views.OutGroupe views.OutDeleteGroupe \
    views.OutUtilisateurs \
    views.InResoudIngredients views.OutResoudIngredients controller.OutLoggin views.InLoggin \
    views.OutIngredientProduits \
    

cd server/models && 
../../../scaneo/scaneo models.go && 
cd ../..