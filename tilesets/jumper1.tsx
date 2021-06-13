<?xml version="1.0" encoding="UTF-8"?>
<tileset version="1.5" tiledversion="1.6.0" name="jumper1" tilewidth="40" tileheight="45" tilecount="11" columns="0">
 <grid orientation="orthogonal" width="1" height="1"/>
 <tile id="0">
  <properties>
   <property name="type" value="collectible"/>
  </properties>
  <image width="40" height="40" source="../images/diamond.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" type="collision" x="6.05147" y="6.05147" width="28.0663" height="27.7533"/>
  </objectgroup>
 </tile>
 <tile id="1">
  <image width="40" height="40" source="../images/gear.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="1.25203" y="1.35636" width="37.0392" height="37.0392"/>
  </objectgroup>
 </tile>
 <tile id="2">
  <image width="40" height="40" source="../images/goal.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="0.208671" y="0.104336" width="39.6476" height="39.9606"/>
  </objectgroup>
 </tile>
 <tile id="3">
  <image width="40" height="40" source="../images/heart.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="1.25203" y="1.25203" width="37.8739" height="37.4565"/>
  </objectgroup>
 </tile>
 <tile id="4">
  <image width="40" height="40" source="../images/key.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="1.35636" y="10.4336" width="37.5609" height="20.7628"/>
  </objectgroup>
 </tile>
 <tile id="5">
  <image width="40" height="40" source="../images/mine.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="1.4607" y="1.14769" width="37.3522" height="37.8739"/>
  </objectgroup>
 </tile>
 <tile id="6">
  <properties>
   <property name="type" value="solid"/>
  </properties>
  <image width="40" height="40" source="../images/solid.png"/>
 </tile>
 <tile id="7">
  <image width="40" height="45" source="../images/tanky.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="2.71273" y="1.04336" width="34.3265" height="40.1693"/>
  </objectgroup>
 </tile>
 <tile id="8">
  <image width="40" height="40" source="../images/montster1.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="3.44308" y="2.71273" width="35.4741" height="37.0392"/>
  </objectgroup>
 </tile>
 <tile id="9">
  <image width="40" height="40" source="../images/montster2.png"/>
  <objectgroup draworder="index" id="2">
   <object id="1" x="0.417343" y="0.208671" width="39.2302" height="39.7519"/>
  </objectgroup>
 </tile>
 <tile id="10">
  <properties>
   <property name="type" value="solid"/>
  </properties>
  <image width="40" height="40" source="../images/door.png"/>
 </tile>
</tileset>
