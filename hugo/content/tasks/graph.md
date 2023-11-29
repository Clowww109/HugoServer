---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

{{< mermaid >}}
graph LR
Margaretta[square] --> Elliott[round-rect]
Darrion[rect] --> Margaretta[square]
Darrion[rect] --> Warren[round-rect]
Darrion[rect] --> Shana[round-rect]
Earnestine[square] --> Margaretta[square]
Earnestine[square] --> Polly[rect]
Earnestine[square] --> Sofia[ellipse]
Earnestine[square] --> Elliott[round-rect]
Maryam[ellipse] --> Margaretta[square]
Maryam[ellipse] --> Emil((circle))
Maryam[ellipse] --> Edgar[rect]
Maryam[ellipse] --> Warren[round-rect]
Maryam[ellipse] --> Elliott[round-rect]
Emil((circle)) --> Margaretta[square]
Coralie[round-rect] --> Margaretta[square]
Percy{rhombus} --> Margaretta[square]
Orville[ellipse] --> Margaretta[square]
Orville[ellipse] --> Alessia[rect]
Orville[ellipse] --> Nathaniel[square]
Orville[ellipse] --> Emmitt{rhombus}
Polly[rect] --> Margaretta[square]
Polly[rect] --> Tyrese[ellipse]
Polly[rect] --> Nona((circle))
Polly[rect] --> Raphaelle[rect]
Edgar[rect] --> Margaretta[square]
Edgar[rect] --> Sofia[ellipse]
Edgar[rect] --> Warren[round-rect]
Corrine{rhombus} --> Margaretta[square]
Corrine{rhombus} --> Tyrese[ellipse]
Alessia[rect] --> Margaretta[square]
Sydni[square] --> Margaretta[square]
Ellie[rect] --> Margaretta[square]
Ellie[rect] --> Raphaelle[rect]
Sofia[ellipse] --> Margaretta[square]
Sofia[ellipse] --> Tyrese[ellipse]
Nathaniel[square] --> Margaretta[square]
Nathaniel[square] --> Adaline[rect]
Nathaniel[square] --> Thora{rhombus}
Nathaniel[square] --> Maudie{rhombus}
Tyrese[ellipse] --> Margaretta[square]
Tyrese[ellipse] --> Maudie{rhombus}
Adaline[rect] --> Margaretta[square]
Adaline[rect] --> Maudie{rhombus}
Emmitt{rhombus} --> Margaretta[square]
Nona((circle)) --> Margaretta[square]
Raphaelle[rect] --> Margaretta[square]
Thora{rhombus} --> Margaretta[square]
Thora{rhombus} --> Elliott[round-rect]
Thora{rhombus} --> Shana[round-rect]
Warren[round-rect] --> Margaretta[square]
Maudie{rhombus} --> Margaretta[square]
Elliott[round-rect] --> Margaretta[square]
Shana[round-rect] --> Margaretta[square]
{{< /mermaid >}}