(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["fournisseurs"],{"4acb":function(e,t,i){"use strict";i.r(t);var n=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("v-dialog",{attrs:{"max-width":"500"}}),i("v-container",[i("liste-fournisseurs")],1)],1)},r=[],s=(i("96cf"),i("1da1")),a=i("d4ec"),o=i("bee2"),c=i("262e"),u=i("2caf"),l=i("9ab4"),d=i("2b0e"),h=i("2fe1"),v=i("0f2e"),p=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("div",[i("v-dialog",{attrs:{"max-width":"500px"},model:{value:e.showEditFournisseur,callback:function(t){e.showEditFournisseur=t},expression:"showEditFournisseur"}},[i("details-fournisseur",{attrs:{fournisseur:e.currentFournisseur,editMode:e.editMode},on:{accept:e.onEditFournisseurDone}})],1),i("v-dialog",{attrs:{"max-width":"500px"},model:{value:e.showConfirmeSupprimeFournisseur,callback:function(t){e.showConfirmeSupprimeFournisseur=t},expression:"showConfirmeSupprimeFournisseur"}},[i("v-card",[i("v-card-title",[e._v("Confirmer la suppression")]),i("v-card-text",[e._v(" Le fournisseur et tout les produits associés seront supprimés. "),i("br"),e._v(" Attention, cette opération est "),i("b",[e._v("irréversible")]),e._v(". ")]),i("v-card-actions",[i("v-spacer"),i("v-btn",{attrs:{color:"error"},on:{click:e.deleteFournisseur}},[e._v("Supprimer définitivement")])],1)],1)],1),i("v-dialog",{attrs:{"max-width":"800px"},model:{value:e.showEditLivraison,callback:function(t){e.showEditLivraison=t},expression:"showEditLivraison"}},[i("details-livraison",{attrs:{livraison:e.currentLivraison,editMode:e.editMode},on:{accept:e.onEditLivraisonDone}})],1),i("v-dialog",{attrs:{"max-width":"500px"},model:{value:e.showConfirmeSupprimeLivraison,callback:function(t){e.showConfirmeSupprimeLivraison=t},expression:"showConfirmeSupprimeLivraison"}},[i("v-card",[i("v-card-title",[e._v("Confirmer la suppression")]),i("v-card-text",[e._v(" La contrainte de livraison sera supprimée et retirée de tous les produits associés. "),i("br"),e._v(" Attention, cette opération est "),i("b",[e._v("irréversible")]),e._v(". ")]),i("v-card-actions",[i("v-spacer"),i("v-btn",{attrs:{color:"error"},on:{click:e.deleteLivraison}},[e._v("Supprimer définitivement")])],1)],1)],1),i("v-toolbar",{attrs:{dense:""}},[i("v-toolbar-title",[e._v("Fournisseurs et contraintes de livraisons")]),i("v-spacer"),i("v-toolbar-items",[i("tooltip-btn",{attrs:{tooltip:"Ajouter un fournisseur...","mdi-icon":"plus",color:"green"},on:{click:function(t){return e.startCreateFournisseur()}}})],1)],1),i("v-treeview",{staticClass:"my-2",attrs:{items:e.treeItems,dense:"","open-on-click":""},scopedSlots:e._u([{key:"label",fn:function(t){var n=t.item;return[e.asTI(n).isFournisseur?i("v-row",{attrs:{"no-gutters":""}},[i("v-col",{staticClass:"align-self-center"},[e._v(" "+e._s(e.asTIF(n).fournisseur.nom)+" - "),i("i",[e._v(e._s(e.asTIF(n).fournisseur.lieu))])]),i("v-col",{staticClass:"align-self-center text-right"},[i("tooltip-btn",{attrs:{tooltip:"Ajouter une contrainte de livraison...","mdi-icon":"plus",color:"green"},on:{click:function(t){e.startCreateLivraison(e.asTIF(n).fournisseur)}}}),i("tooltip-btn",{attrs:{"mdi-icon":"pencil",tooltip:"Modifier ce fournisseur...",color:"secondary"},on:{click:function(t){e.startEditFournisseur(e.asTIF(n).fournisseur)}}}),i("tooltip-btn",{attrs:{"mdi-icon":"close",color:"red",tooltip:"Supprimer le fournisseur et les produits associés"},on:{click:function(t){e.confirmeDeleteFournisseur(e.asTIF(n).fournisseur)}}})],1)],1):i("v-row",{attrs:{"no-gutters":""}},[i("v-col",{staticClass:"align-self-center",attrs:{cols:"2"}},[i("span",{domProps:{innerHTML:e._s(e.formatLivraisonNom(e.asTIL(n).livraison))}})]),i("v-col",{staticClass:"align-self-center",attrs:{cols:"8"}},e._l(e.filterJoursLivraison(e.asTIL(n).livraison),(function(t){return i("v-chip",{key:t,attrs:{small:""}},[e._v(" "+e._s(t)+" ")])})),1),i("v-col",{staticClass:"align-self-center text-right"},[i("tooltip-btn",{attrs:{"mdi-icon":"pencil",tooltip:"Modifier cette contrainte de livraison...",color:"secondary"},on:{click:function(t){e.startEditLivraison(e.asTIL(n).livraison)}}}),i("tooltip-btn",{attrs:{"mdi-icon":"close",color:"red",tooltip:"Supprimer la contrainte de livraison..."},on:{click:function(t){e.confirmeDeleteLivraison(e.asTIL(n).livraison)}}})],1)],1)]}}])})],1)},f=[],m=(i("4de4"),i("4160"),i("d81d"),i("07ac"),i("159b"),i("1c9b")),b=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("v-card",[i("v-card-title",{attrs:{"primary-title":""}},[e._v(" Détails du fournisseur ")]),i("v-card-text",[i("v-form",[i("v-text-field",{attrs:{label:"Nom du fournisseur",required:""},model:{value:e.innerFournisseur.nom,callback:function(t){e.$set(e.innerFournisseur,"nom",t)},expression:"innerFournisseur.nom"}}),i("v-text-field",{attrs:{label:"Lieu",hint:"Le lieu permet de sélectionner rapidement un groupe de fournisseurs.","persistent-hint":""},model:{value:e.innerFournisseur.lieu,callback:function(t){e.$set(e.innerFournisseur,"lieu",t)},expression:"innerFournisseur.lieu"}})],1)],1),i("v-card-actions",[i("v-spacer"),i("v-btn",{attrs:{color:"success"},on:{click:function(t){return e.$emit("accept",e.innerFournisseur)}}},[e._v(" "+e._s("new"==e.editMode?"Créer":"Enregistrer")+" ")])],1)],1)},g=[],O=i("5dc2"),C=i("60a3"),j=d["a"].extend({props:{fournisseur:Object,editMode:String}}),w=function(e){Object(c["a"])(i,e);var t=Object(u["a"])(i);function i(){var e;return Object(a["a"])(this,i),e=t.apply(this,arguments),e.innerFournisseur=e.duplique(),e}return Object(o["a"])(i,[{key:"_",value:function(){this.innerFournisseur=this.duplique()}},{key:"duplique",value:function(){return null==this.fournisseur?{nom:"",lieu:""}:Object(O["a"])(this.fournisseur)}}]),i}(j);Object(l["a"])([Object(C["a"])("fournisseur")],w.prototype,"_",null),w=Object(l["a"])([Object(h["b"])({})],w);var y=w,S=y,x=i("2877"),L=i("6544"),k=i.n(L),_=i("8336"),F=i("b0af"),I=i("99d9"),E=i("4bd4"),T=i("2fa4"),A=i("8654"),V=Object(x["a"])(S,b,g,!1,null,"91819830",null),$=V.exports;k()(V,{VBtn:_["a"],VCard:F["a"],VCardActions:I["a"],VCardText:I["c"],VCardTitle:I["d"],VForm:E["a"],VSpacer:T["a"],VTextField:A["a"]});var D=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("v-card",[i("v-card-title",{attrs:{"primary-title":""}},[e._v(" Détails de la contrainte de livraison ")]),i("v-card-text",[i("v-form",[i("v-row",[i("v-col",{attrs:{cols:"5"}},[i("v-text-field",{attrs:{label:"Label de la contrainte de livraison",required:""},model:{value:e.innerLivraison.nom,callback:function(t){e.$set(e.innerLivraison,"nom",t)},expression:"innerLivraison.nom"}}),i("v-select",{attrs:{items:e.optionsFournisseurs,rules:[e.rules.idRequired],label:"Fournisseur"},model:{value:e.innerLivraison.id_fournisseur,callback:function(t){e.$set(e.innerLivraison,"id_fournisseur",t)},expression:"innerLivraison.id_fournisseur"}})],1),i("v-col",[i("jours-livraison-field",{model:{value:e.innerLivraison.jours_livraison,callback:function(t){e.$set(e.innerLivraison,"jours_livraison",t)},expression:"innerLivraison.jours_livraison"}}),i("v-text-field",{attrs:{label:"Delai de livraison",type:"number",hint:"La date de commande est avancé de ce nombre de jours ouvrés par rapport à la date de livraison."},model:{value:e.innerLivraison.delai_commande,callback:function(t){e.$set(e.innerLivraison,"delai_commande",e._n(t))},expression:"innerLivraison.delai_commande"}}),i("v-text-field",{attrs:{label:"Anticipation",type:"number",hint:"La date de livraison est avancé de ce nombre de jours par rapport à la date d'utilisation."},model:{value:e.innerLivraison.anticipation,callback:function(t){e.$set(e.innerLivraison,"anticipation",e._n(t))},expression:"innerLivraison.anticipation"}})],1)],1)],1)],1),i("v-card-actions",[i("v-spacer"),i("v-btn",{attrs:{color:"success"},on:{click:function(t){return e.$emit("accept",e.innerLivraison)}}},[e._v(" "+e._s("new"==e.editMode?"Créer":"Enregistrer")+" ")])],1)],1)},M=[],P=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("v-select",{attrs:{multiple:"",chips:"",items:e.items,label:"Jours de livraison",hint:"Jours possible de livraison."},model:{value:e.innerJours,callback:function(t){e.innerJours=t},expression:"innerJours"}})},K=[],B=i("fe36"),R=d["a"].extend({props:{jours:Array},model:{prop:"jours",event:"change"}}),N=function(e){Object(c["a"])(i,e);var t=Object(u["a"])(i);function i(){var e;return Object(a["a"])(this,i),e=t.apply(this,arguments),e.items=B["a"].map((function(e,t){return{text:e,value:t}})),e}return Object(o["a"])(i,[{key:"innerJours",get:function(){var e=[];return this.jours.forEach((function(t,i){t&&e.push(i)})),e},set:function(e){var t=this.items.map((function(e){return!1}));e.forEach((function(e){return t[e]=!0})),this.$emit("change",t)}}]),i}(R);N=Object(l["a"])([Object(h["b"])({})],N);var q=N,J=q,W=i("b974"),H=Object(x["a"])(J,P,K,!1,null,"9dfa8e9e",null),z=H.exports;k()(H,{VSelect:W["a"]});var G=d["a"].extend({props:{livraison:Object,editMode:String}}),Q=function(e){Object(c["a"])(i,e);var t=Object(u["a"])(i);function i(){var e;return Object(a["a"])(this,i),e=t.apply(this,arguments),e.innerLivraison=e.duplique(),e.rules={idRequired:function(e){return e>=0||"Champ requis"}},e}return Object(o["a"])(i,[{key:"_",value:function(){this.innerLivraison=this.duplique()}},{key:"duplique",value:function(){return null==this.livraison?Object(O["b"])():Object(O["a"])(this.livraison)}},{key:"optionsFournisseurs",get:function(){if(null==v["a"].data)return[];var e=Object.values(v["a"].data.fournisseurs||{}).map((function(e){return{text:e.nom,value:e.id}}));return e}}]),i}(G);Object(l["a"])([Object(C["a"])("livraison")],Q.prototype,"_",null),Q=Object(l["a"])([Object(h["b"])({components:{JoursLivraisonField:z}})],Q);var U=Q,X=U,Y=i("62ad"),Z=i("0fd9"),ee=Object(x["a"])(X,D,M,!1,null,"20e8a986",null),te=ee.exports;k()(ee,{VBtn:_["a"],VCard:F["a"],VCardActions:I["a"],VCardText:I["c"],VCardTitle:I["d"],VCol:Y["a"],VForm:E["a"],VRow:Z["a"],VSelect:W["a"],VSpacer:T["a"],VTextField:A["a"]});var ie=d["a"].extend({props:{}}),ne=function(e){Object(c["a"])(i,e);var t=Object(u["a"])(i);function i(){var e;return Object(a["a"])(this,i),e=t.apply(this,arguments),e.showConfirmeSupprimeFournisseur=!1,e.showEditFournisseur=!1,e.currentFournisseur=null,e.editMode="new",e.showConfirmeSupprimeLivraison=!1,e.showEditLivraison=!1,e.currentLivraison=null,e.asTI=function(e){return e},e.asTIF=function(e){return e},e.asTIL=function(e){return e},e}return Object(o["a"])(i,[{key:"startCreateFournisseur",value:function(){this.currentFournisseur={nom:"",lieu:""},this.editMode="new",this.showEditFournisseur=!0}},{key:"startEditFournisseur",value:function(e){this.currentFournisseur=e,this.editMode="edit",this.showEditFournisseur=!0}},{key:"onEditFournisseurDone",value:function(){var e=Object(s["a"])(regeneratorRuntime.mark((function e(t){var i;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(this.showEditFournisseur=!1,"new"!=this.editMode){e.next=7;break}return e.next=4,v["a"].data.createFournisseur(t);case 4:i="Fournisseur ajouté avec succès.",e.next=10;break;case 7:return e.next=9,v["a"].data.updateFournisseur(t);case 9:i="Fournisseur édité avec succès.";case 10:null==v["a"].notifications.getError()&&v["a"].notifications.setMessage(i);case 11:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"confirmeDeleteFournisseur",value:function(e){this.currentFournisseur=e,this.showConfirmeSupprimeFournisseur=!0}},{key:"deleteFournisseur",value:function(){var e=Object(s["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(null!=this.currentFournisseur&&void 0!=this.currentFournisseur.id){e.next=2;break}return e.abrupt("return");case 2:return this.showConfirmeSupprimeFournisseur=!1,e.next=5,v["a"].data.deleteFournisseur(this.currentFournisseur.id);case 5:return e.next=7,v["a"].data.loadSejours();case 7:null==v["a"].notifications.getError()&&v["a"].notifications.setMessage("Fournisseur supprimé avec succès.");case 8:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"formatFournisseur",value:function(e){return v["a"].getFournisseur(e).nom}},{key:"formatLivraisonNom",value:function(e){var t=e.nom;return""==t?"<i class='grey--text'>par défaut</i>":t}},{key:"filterJoursLivraison",value:function(e){var t=[];return e.jours_livraison.forEach((function(e,i){e&&t.push(B["a"][i])})),t}},{key:"startCreateLivraison",value:function(e){this.currentLivraison=Object(O["b"])(),this.currentLivraison.id_fournisseur=e.id,this.editMode="new",this.showEditLivraison=!0}},{key:"startEditLivraison",value:function(e){this.currentLivraison=e,this.editMode="edit",this.showEditLivraison=!0}},{key:"onEditLivraisonDone",value:function(){var e=Object(s["a"])(regeneratorRuntime.mark((function e(t){var i;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(this.showEditLivraison=!1,"new"!=this.editMode){e.next=7;break}return e.next=4,v["a"].data.createLivraison(t);case 4:i="Contrainte de livraison ajoutée avec succès.",e.next=10;break;case 7:return e.next=9,v["a"].data.updateLivraison(t);case 9:i="Contrainte de livraison éditée avec succès.";case 10:null==v["a"].notifications.getError()&&v["a"].notifications.setMessage(i);case 11:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}()},{key:"confirmeDeleteLivraison",value:function(e){this.currentLivraison=e,this.showConfirmeSupprimeLivraison=!0}},{key:"deleteLivraison",value:function(){var e=Object(s["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(null!=this.currentLivraison&&void 0!=this.currentLivraison.id){e.next=2;break}return e.abrupt("return");case 2:return this.showConfirmeSupprimeLivraison=!1,e.next=5,v["a"].data.deleteLivraison(this.currentLivraison.id);case 5:null==v["a"].notifications.getError()&&v["a"].notifications.setMessage("Contrainte de livraison supprimée avec succès.");case 6:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}()},{key:"treeItems",get:function(){var e=Object.values(v["a"].data.fournisseurs||{}),t=Object.values(v["a"].data.livraisons||{});return e.map((function(e){var i=t.filter((function(t){return t.id_fournisseur==e.id})),n=i.map((function(e){return{isFournisseur:!1,livraison:e}}));return{fournisseur:e,children:n,isFournisseur:!0}}))}}]),i}(ie);ne=Object(l["a"])([Object(h["b"])({components:{TooltipBtn:m["a"],DetailsFournisseur:$,DetailsLivraison:te}})],ne);var re=ne,se=re,ae=(i("6f0e"),i("cc20")),oe=i("169a"),ce=i("71d9"),ue=i("2a7f"),le=(i("13d5"),i("4ec9"),i("b64b"),i("d3b7"),i("ac1f"),i("6062"),i("3ca3"),i("841c"),i("ddb0"),i("3835")),de=i("b85c"),he=i("2909"),ve=i("5530"),pe=(i("fa9e"),i("caad"),i("a9e3"),i("ade3")),fe=i("0789"),me=i("132d"),be=i("3206"),ge=i("a9ad"),Oe=i("58df"),Ce=i("80d2"),je=Object(Oe["a"])(ge["a"],Object(be["a"])("treeview")),we={activatable:Boolean,activeClass:{type:String,default:"v-treeview-node--active"},color:{type:String,default:"primary"},expandIcon:{type:String,default:"$subgroup"},indeterminateIcon:{type:String,default:"$checkboxIndeterminate"},itemChildren:{type:String,default:"children"},itemDisabled:{type:String,default:"disabled"},itemKey:{type:String,default:"id"},itemText:{type:String,default:"name"},loadChildren:Function,loadingIcon:{type:String,default:"$loading"},offIcon:{type:String,default:"$checkboxOff"},onIcon:{type:String,default:"$checkboxOn"},openOnClick:Boolean,rounded:Boolean,selectable:Boolean,selectedColor:{type:String,default:"accent"},shaped:Boolean,transition:Boolean,selectionType:{type:String,default:"leaf",validator:function(e){return["leaf","independent"].includes(e)}}},ye=je.extend().extend({name:"v-treeview-node",inject:{treeview:{default:null}},props:Object(ve["a"])({level:Number,item:{type:Object,default:function(){return null}},parentIsDisabled:Boolean},we),data:function(){return{hasLoaded:!1,isActive:!1,isIndeterminate:!1,isLoading:!1,isOpen:!1,isSelected:!1}},computed:{disabled:function(){return Object(Ce["n"])(this.item,this.itemDisabled)||this.parentIsDisabled&&"leaf"===this.selectionType},key:function(){return Object(Ce["n"])(this.item,this.itemKey)},children:function(){return Object(Ce["n"])(this.item,this.itemChildren)},text:function(){return Object(Ce["n"])(this.item,this.itemText)},scopedProps:function(){return{item:this.item,leaf:!this.children,selected:this.isSelected,indeterminate:this.isIndeterminate,active:this.isActive,open:this.isOpen}},computedIcon:function(){return this.isIndeterminate?this.indeterminateIcon:this.isSelected?this.onIcon:this.offIcon},hasChildren:function(){return!!this.children&&(!!this.children.length||!!this.loadChildren)}},created:function(){this.treeview.register(this)},beforeDestroy:function(){this.treeview.unregister(this)},methods:{checkChildren:function(){var e=this;return new Promise((function(t){if(!e.children||e.children.length||!e.loadChildren||e.hasLoaded)return t();e.isLoading=!0,t(e.loadChildren(e.item))})).then((function(){e.isLoading=!1,e.hasLoaded=!0}))},open:function(){this.isOpen=!this.isOpen,this.treeview.updateOpen(this.key,this.isOpen),this.treeview.emitOpen()},genLabel:function(){var e=[];return this.$scopedSlots.label?e.push(this.$scopedSlots.label(this.scopedProps)):e.push(this.text),this.$createElement("div",{slot:"label",staticClass:"v-treeview-node__label"},e)},genPrependSlot:function(){return this.$scopedSlots.prepend?this.$createElement("div",{staticClass:"v-treeview-node__prepend"},this.$scopedSlots.prepend(this.scopedProps)):null},genAppendSlot:function(){return this.$scopedSlots.append?this.$createElement("div",{staticClass:"v-treeview-node__append"},this.$scopedSlots.append(this.scopedProps)):null},genContent:function(){var e=[this.genPrependSlot(),this.genLabel(),this.genAppendSlot()];return this.$createElement("div",{staticClass:"v-treeview-node__content"},e)},genToggle:function(){var e=this;return this.$createElement(me["a"],{staticClass:"v-treeview-node__toggle",class:{"v-treeview-node__toggle--open":this.isOpen,"v-treeview-node__toggle--loading":this.isLoading},slot:"prepend",on:{click:function(t){t.stopPropagation(),e.isLoading||e.checkChildren().then((function(){return e.open()}))}}},[this.isLoading?this.loadingIcon:this.expandIcon])},genCheckbox:function(){var e=this;return this.$createElement(me["a"],{staticClass:"v-treeview-node__checkbox",props:{color:this.isSelected||this.isIndeterminate?this.selectedColor:void 0,disabled:this.disabled},on:{click:function(t){t.stopPropagation(),e.isLoading||e.checkChildren().then((function(){e.$nextTick((function(){e.isSelected=!e.isSelected,e.isIndeterminate=!1,e.treeview.updateSelected(e.key,e.isSelected),e.treeview.emitSelected()}))}))}}},[this.computedIcon])},genLevel:function(e){var t=this;return Object(Ce["h"])(e).map((function(){return t.$createElement("div",{staticClass:"v-treeview-node__level"})}))},genNode:function(){var e=this,t=[this.genContent()];return this.selectable&&t.unshift(this.genCheckbox()),this.hasChildren?t.unshift(this.genToggle()):t.unshift.apply(t,Object(he["a"])(this.genLevel(1))),t.unshift.apply(t,Object(he["a"])(this.genLevel(this.level))),this.$createElement("div",this.setTextColor(this.isActive&&this.color,{staticClass:"v-treeview-node__root",class:Object(pe["a"])({},this.activeClass,this.isActive),on:{click:function(){e.openOnClick&&e.hasChildren?e.checkChildren().then(e.open):e.activatable&&!e.disabled&&(e.isActive=!e.isActive,e.treeview.updateActive(e.key,e.isActive),e.treeview.emitActive())}}}),t)},genChild:function(e,t){return this.$createElement(ye,{key:Object(Ce["n"])(e,this.itemKey),props:{activatable:this.activatable,activeClass:this.activeClass,item:e,selectable:this.selectable,selectedColor:this.selectedColor,color:this.color,expandIcon:this.expandIcon,indeterminateIcon:this.indeterminateIcon,offIcon:this.offIcon,onIcon:this.onIcon,loadingIcon:this.loadingIcon,itemKey:this.itemKey,itemText:this.itemText,itemDisabled:this.itemDisabled,itemChildren:this.itemChildren,loadChildren:this.loadChildren,transition:this.transition,openOnClick:this.openOnClick,rounded:this.rounded,shaped:this.shaped,level:this.level+1,selectionType:this.selectionType,parentIsDisabled:t},scopedSlots:this.$scopedSlots})},genChildrenWrapper:function(){var e=this;if(!this.isOpen||!this.children)return null;var t=[this.children.map((function(t){return e.genChild(t,e.disabled)}))];return this.$createElement("div",{staticClass:"v-treeview-node__children"},t)},genTransition:function(){return this.$createElement(fe["a"],[this.genChildrenWrapper()])}},render:function(e){var t=[this.genNode()];return this.transition?t.push(this.genTransition()):t.push(this.genChildrenWrapper()),e("div",{staticClass:"v-treeview-node",class:{"v-treeview-node--leaf":!this.hasChildren,"v-treeview-node--click":this.openOnClick,"v-treeview-node--disabled":this.disabled,"v-treeview-node--rounded":this.rounded,"v-treeview-node--shaped":this.shaped,"v-treeview-node--selected":this.isSelected,"v-treeview-node--excluded":this.treeview.isExcluded(this.key)},attrs:{"aria-expanded":String(this.isOpen)}},t)}}),Se=ye,xe=i("7560"),Le=i("d9bd");i("c975");function ke(e,t,i){var n=Object(Ce["n"])(e,i);return n.toLocaleLowerCase().indexOf(t.toLocaleLowerCase())>-1}function _e(e,t,i,n,r,s,a){if(e(t,i,r))return!0;var o=Object(Ce["n"])(t,s);if(o){for(var c=!1,u=0;u<o.length;u++)_e(e,o[u],i,n,r,s,a)&&(c=!0);if(c)return!0}return a.add(Object(Ce["n"])(t,n)),!1}var Fe=Object(Oe["a"])(Object(be["b"])("treeview"),xe["a"]).extend({name:"v-treeview",provide:function(){return{treeview:this}},props:Object(ve["a"])({active:{type:Array,default:function(){return[]}},dense:Boolean,filter:Function,hoverable:Boolean,items:{type:Array,default:function(){return[]}},multipleActive:Boolean,open:{type:Array,default:function(){return[]}},openAll:Boolean,returnObject:{type:Boolean,default:!1},search:String,value:{type:Array,default:function(){return[]}}},we),data:function(){return{level:-1,activeCache:new Set,nodes:{},openCache:new Set,selectedCache:new Set}},computed:{excludedItems:function(){var e=new Set;if(!this.search)return e;for(var t=0;t<this.items.length;t++)_e(this.filter||ke,this.items[t],this.search,this.itemKey,this.itemText,this.itemChildren,e);return e}},watch:{items:{handler:function(){var e=this,t=Object.keys(this.nodes).map((function(t){return Object(Ce["n"])(e.nodes[t].item,e.itemKey)})),i=this.getKeys(this.items),n=Object(Ce["c"])(i,t);if(n.length||!(i.length<t.length)){n.forEach((function(t){return delete e.nodes[t]}));var r=Object(he["a"])(this.selectedCache);this.selectedCache=new Set,this.activeCache=new Set,this.openCache=new Set,this.buildTree(this.items),Object(Ce["j"])(r,Object(he["a"])(this.selectedCache))||this.emitSelected()}},deep:!0},active:function(e){this.handleNodeCacheWatcher(e,this.activeCache,this.updateActive,this.emitActive)},value:function(e){this.handleNodeCacheWatcher(e,this.selectedCache,this.updateSelected,this.emitSelected)},open:function(e){this.handleNodeCacheWatcher(e,this.openCache,this.updateOpen,this.emitOpen)}},created:function(){var e=this,t=function(t){return e.returnObject?Object(Ce["n"])(t,e.itemKey):t};this.buildTree(this.items);var i,n=Object(de["a"])(this.value.map(t));try{for(n.s();!(i=n.n()).done;){var r=i.value;this.updateSelected(r,!0,!0)}}catch(c){n.e(c)}finally{n.f()}var s,a=Object(de["a"])(this.active.map(t));try{for(a.s();!(s=a.n()).done;){var o=s.value;this.updateActive(o,!0)}}catch(c){a.e(c)}finally{a.f()}},mounted:function(){var e=this;(this.$slots.prepend||this.$slots.append)&&Object(Le["c"])("The prepend and append slots require a slot-scope attribute",this),this.openAll?this.updateAll(!0):(this.open.forEach((function(t){return e.updateOpen(e.returnObject?Object(Ce["n"])(t,e.itemKey):t,!0)})),this.emitOpen())},methods:{updateAll:function(e){var t=this;Object.keys(this.nodes).forEach((function(i){return t.updateOpen(Object(Ce["n"])(t.nodes[i].item,t.itemKey),e)})),this.emitOpen()},getKeys:function(e){for(var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[],i=0;i<e.length;i++){var n=Object(Ce["n"])(e[i],this.itemKey);t.push(n);var r=Object(Ce["n"])(e[i],this.itemChildren);r&&t.push.apply(t,Object(he["a"])(this.getKeys(r)))}return t},buildTree:function(e){for(var t=this,i=arguments.length>1&&void 0!==arguments[1]?arguments[1]:null,n=0;n<e.length;n++){var r=e[n],s=Object(Ce["n"])(r,this.itemKey),a=Object(Ce["n"])(r,this.itemChildren,[]),o=this.nodes.hasOwnProperty(s)?this.nodes[s]:{isSelected:!1,isIndeterminate:!1,isActive:!1,isOpen:!1,vnode:null},c={vnode:o.vnode,parent:i,children:a.map((function(e){return Object(Ce["n"])(e,t.itemKey)})),item:r};if(this.buildTree(a,s),!this.nodes.hasOwnProperty(s)&&null!==i&&this.nodes.hasOwnProperty(i)?c.isSelected=this.nodes[i].isSelected:(c.isSelected=o.isSelected,c.isIndeterminate=o.isIndeterminate),c.isActive=o.isActive,c.isOpen=o.isOpen,this.nodes[s]=c,a.length){var u=this.calculateState(s,this.nodes),l=u.isSelected,d=u.isIndeterminate;c.isSelected=l,c.isIndeterminate=d}!this.nodes[s].isSelected||"independent"!==this.selectionType&&0!==c.children.length||this.selectedCache.add(s),this.nodes[s].isActive&&this.activeCache.add(s),this.nodes[s].isOpen&&this.openCache.add(s),this.updateVnodeState(s)}},calculateState:function(e,t){var i=t[e].children,n=i.reduce((function(e,i){return e[0]+=+Boolean(t[i].isSelected),e[1]+=+Boolean(t[i].isIndeterminate),e}),[0,0]),r=!!i.length&&n[0]===i.length,s=!r&&(n[0]>0||n[1]>0);return{isSelected:r,isIndeterminate:s}},emitOpen:function(){this.emitNodeCache("update:open",this.openCache)},emitSelected:function(){this.emitNodeCache("input",this.selectedCache)},emitActive:function(){this.emitNodeCache("update:active",this.activeCache)},emitNodeCache:function(e,t){var i=this;this.$emit(e,this.returnObject?Object(he["a"])(t).map((function(e){return i.nodes[e].item})):Object(he["a"])(t))},handleNodeCacheWatcher:function(e,t,i,n){var r=this;e=this.returnObject?e.map((function(e){return Object(Ce["n"])(e,r.itemKey)})):e;var s=Object(he["a"])(t);Object(Ce["j"])(s,e)||(s.forEach((function(e){return i(e,!1)})),e.forEach((function(e){return i(e,!0)})),n())},getDescendants:function(e){var t,i=arguments.length>1&&void 0!==arguments[1]?arguments[1]:[],n=this.nodes[e].children;(t=i).push.apply(t,Object(he["a"])(n));for(var r=0;r<n.length;r++)i=this.getDescendants(n[r],i);return i},getParents:function(e){var t=this.nodes[e].parent,i=[];while(null!==t)i.push(t),t=this.nodes[t].parent;return i},register:function(e){var t=Object(Ce["n"])(e.item,this.itemKey);this.nodes[t].vnode=e,this.updateVnodeState(t)},unregister:function(e){var t=Object(Ce["n"])(e.item,this.itemKey);this.nodes[t]&&(this.nodes[t].vnode=null)},isParent:function(e){return this.nodes[e].children&&this.nodes[e].children.length},updateActive:function(e,t){var i=this;if(this.nodes.hasOwnProperty(e)){this.multipleActive||this.activeCache.forEach((function(e){i.nodes[e].isActive=!1,i.updateVnodeState(e),i.activeCache.delete(e)}));var n=this.nodes[e];n&&(t?this.activeCache.add(e):this.activeCache.delete(e),n.isActive=t,this.updateVnodeState(e))}},updateSelected:function(e,t){var i=arguments.length>2&&void 0!==arguments[2]&&arguments[2];if(this.nodes.hasOwnProperty(e)){var n=new Map;if("independent"!==this.selectionType){var r,s=Object(de["a"])(this.getDescendants(e));try{for(s.s();!(r=s.n()).done;){var a=r.value;Object(Ce["n"])(this.nodes[a].item,this.itemDisabled)&&!i||(this.nodes[a].isSelected=t,this.nodes[a].isIndeterminate=!1,n.set(a,t))}}catch(b){s.e(b)}finally{s.f()}var o=this.calculateState(e,this.nodes);this.nodes[e].isSelected=t,this.nodes[e].isIndeterminate=o.isIndeterminate,n.set(e,t);var c,u=Object(de["a"])(this.getParents(e));try{for(u.s();!(c=u.n()).done;){var l=c.value,d=this.calculateState(l,this.nodes);this.nodes[l].isSelected=d.isSelected,this.nodes[l].isIndeterminate=d.isIndeterminate,n.set(l,d.isSelected)}}catch(b){u.e(b)}finally{u.f()}}else this.nodes[e].isSelected=t,this.nodes[e].isIndeterminate=!1,n.set(e,t);var h,v=Object(de["a"])(n.entries());try{for(v.s();!(h=v.n()).done;){var p=Object(le["a"])(h.value,2),f=p[0],m=p[1];this.updateVnodeState(f),"leaf"===this.selectionType&&this.isParent(f)||(!0===m?this.selectedCache.add(f):this.selectedCache.delete(f))}}catch(b){v.e(b)}finally{v.f()}}},updateOpen:function(e,t){var i=this;if(this.nodes.hasOwnProperty(e)){var n=this.nodes[e],r=Object(Ce["n"])(n.item,this.itemChildren);r&&!r.length&&n.vnode&&!n.vnode.hasLoaded?n.vnode.checkChildren().then((function(){return i.updateOpen(e,t)})):r&&r.length&&(n.isOpen=t,n.isOpen?this.openCache.add(e):this.openCache.delete(e),this.updateVnodeState(e))}},updateVnodeState:function(e){var t=this.nodes[e];t&&t.vnode&&(t.vnode.isSelected=t.isSelected,t.vnode.isIndeterminate=t.isIndeterminate,t.vnode.isActive=t.isActive,t.vnode.isOpen=t.isOpen)},isExcluded:function(e){return!!this.search&&this.excludedItems.has(e)}},render:function(e){var t=this,i=this.items.length?this.items.map((function(e){var i=Se.options.methods.genChild.bind(t);return i(e,Object(Ce["n"])(e,t.itemDisabled))})):this.$slots.default;return e("div",{staticClass:"v-treeview",class:Object(ve["a"])({"v-treeview--hoverable":this.hoverable,"v-treeview--dense":this.dense},this.themeClasses)},i)}}),Ie=Object(x["a"])(se,p,f,!1,null,null,null),Ee=Ie.exports;k()(Ie,{VBtn:_["a"],VCard:F["a"],VCardActions:I["a"],VCardText:I["c"],VCardTitle:I["d"],VChip:ae["a"],VCol:Y["a"],VDialog:oe["a"],VRow:Z["a"],VSpacer:T["a"],VToolbar:ce["a"],VToolbarItems:ue["a"],VToolbarTitle:ue["b"],VTreeview:Fe});var Te=d["a"].extend({props:{}}),Ae=function(e){Object(c["a"])(i,e);var t=Object(u["a"])(i);function i(){return Object(a["a"])(this,i),t.apply(this,arguments)}return Object(o["a"])(i,[{key:"mounted",value:function(){var e=Object(s["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,v["a"].data.loadFournisseurs();case 2:null==v["a"].notifications.getError()&&v["a"].notifications.setMessage("Fournisseurs chargés.");case 3:case"end":return e.stop()}}),e)})));function t(){return e.apply(this,arguments)}return t}()}]),i}(Te);Ae=Object(l["a"])([Object(h["b"])({components:{ListeFournisseurs:Ee}})],Ae);var Ve=Ae,$e=Ve,De=i("a523"),Me=Object(x["a"])($e,n,r,!1,null,"fd90a27e",null);t["default"]=Me.exports;k()(Me,{VContainer:De["a"],VDialog:oe["a"]})},"6f0e":function(e,t,i){"use strict";var n=i("d67b"),r=i.n(n);r.a},"841c":function(e,t,i){"use strict";var n=i("d784"),r=i("825a"),s=i("1d80"),a=i("129f"),o=i("14c3");n("search",1,(function(e,t,i){return[function(t){var i=s(this),n=void 0==t?void 0:t[e];return void 0!==n?n.call(t,i):new RegExp(t)[e](String(i))},function(e){var n=i(t,e,this);if(n.done)return n.value;var s=r(e),c=String(this),u=s.lastIndex;a(u,0)||(s.lastIndex=0);var l=o(s,c);return a(s.lastIndex,u)||(s.lastIndex=u),null===l?-1:l.index}]}))},d67b:function(e,t,i){},fa9e:function(e,t,i){}}]);
//# sourceMappingURL=fournisseurs.285b483b.js.map