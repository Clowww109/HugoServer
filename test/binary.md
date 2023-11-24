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
Omari[ellipse] --> Mariah[round-rect]
Joel[square] --> Omari[ellipse]
Joel[square] --> Mariah[round-rect]
Ervin[rect] --> Omari[ellipse]
Selmer((circle)) --> Omari[ellipse]
Selmer((circle)) --> Sebastian[rect]
Mariah[round-rect] --> Omari[ellipse]
Sebastian[rect] --> Omari[ellipse]
{{< /mermaid >}}