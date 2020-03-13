cd server/models &&
    ../../../scaneo/scaneo models.go &&
    cd ../..

../struct2ts/main -H -i -D -o frontend/src/logic/types.ts \
    views.OutIngredient views.OutIngredients \
    views.OutRecette views.OutRecettes \
    views.OutMenu views.OutMenus \
    views.OutUtilisateurs \
    views.OutSejour \
    views.OutGroupe \
    views.OutDeleteGroupe \
    views.OutSejours \
    views.InResoudIngredients views.OutResoudIngredients controller.OutLoggin views.InLoggin \
    views.InAssistantCreateRepass \
    views.OutIngredientProduits

go run macros/enums.go >frontend/src/logic/enums.ts
