---
title: "Agents: The Model Is an Indispensable Part — the Rest Is Engineering"
description: "Before anything else: what a model is, what a harness is, and where an agent system actually differs from ordinary software."
pubDate: 2026-11-02
tags: [learning, tech, agent, harness]
series: how-agents-work
series_order: 1
lang: en
authorship: human-ai
draft: false
---

## What exactly are we talking about

The word "agent" has been stretched thin these past two years. Within three sentences, models, frameworks, agents, and harnesses are usually muddled into one pot. Before getting into mechanisms, separate three things.

**The model is a stateless inference function.** Tokens in, tokens out, with probability deciding which token comes next; the weights do not change because of a conversation — feed it the same input and it faces the same distribution, nothing carried over from the previous call. It owns no files and no processes. The "it remembers me" you see in products — cross-session memory, chat history — none of it actually lives in the model. The shell stores it and splices it back into the input before each call; chapter 4 goes into this in depth. By itself the model is not an agent — just a component, and one that is occasionally wrong with complete confidence.

**The harness is all the engineering wrapped around the model.** Who decides the next move; what goes into the context and what stays out; how tools are invoked; where commands run; who approves dangerous actions; how history is stored and a session resumed; what the interface looks like. Every behavior that makes the model look like an "agent" lives in this shell. Almost everything in this series is about the shell.

**The environment is the world the agent acts on.** Filesystems, shells, git repositories, browsers, networks, other services — and people.

<figure class="diagram-embed" data-slug="harness-scope" data-aspect="2.00">
<figcaption>Figure 1 · Model, harness, environment, and the three boundaries between them</figcaption>
</figure>

There are three boundaries between these three things.

First, to be clear: these are not official terminology. They are an analytical frame I use to read these systems. The dividing principle is simple: **the model is locked inside the shell and can touch nothing on its own; to affect anything it must pass through an exit, and on the other side there are exactly three kinds of objects, each with a different trust semantics.** Cutting along trust semantics that cannot be merged further yields three seams:

- **The inference boundary**, harness to model: one token round-trip. Across it is the model — the only object whose output is probabilistic, so you cannot assume what comes back. Every uncertainty in the system enters through this one seam; retries, fallbacks, and budgets all orbit it.
- **The side-effect boundary**, harness to environment: where tools actually act. Across it are files, networks, and other services — deterministic, mutable, and often irreversible. The world outside the process changes only past this boundary, which is why sandboxes and approvals guard it.
- **The human boundary**, between person and harness. Across it is a person — who holds ultimate authority and can interrupt at any time: instructions, approvals, interruption. Permission can be delegated; authority never fully is.

Three is the minimal cut that holds: fold the model into "environment" and you lose the key fact that there is only one probabilistic seam; split files and networks further and you add no new trust semantics. Every later chapter traces one of these seams. Real flows cross several at once (a tool call is the model choosing and the harness executing) — the boundaries exist to make the system legible, not to carve it into unrelated pieces.

What is not covered, stated plainly: model training, weights, inference kernels, GPU scheduling — that is the model team's world, and this series does not touch it. Nor is it a collection of prompt tricks. It answers one question only: **what engineering does it take to turn a model into a system that does real work?**

## Where it differs from ordinary software

I take one sentence as the thesis of the whole series: 90% of a harness is ordinary software engineering. The other 10% is not "harder engineering" — it is the chain reaction set off when the component in charge of "what happens next" changes from a deterministic program to a probabilistic model.

An abstract claim is weak, so take an industrial coding agent apart. The next diagram maps the module layout of OpenAI's open-source codex — written in Rust, a hundred-plus crates, one of the most complete specimens available today.

<figure class="diagram-embed" data-slug="harness-map" data-aspect="2.00">
<figcaption>Figure 2 · Module map of an industrial harness (after codex's crate layout): four pillars — grey, orange, red, cyan — classified by how they relate to the model</figcaption>
</figure>

Four pillars:

- **Grey: the 90% floor.** Interface, network proxy, secrets, files and git, installers. These modules exist with or without an LLM; you build them the way you always have, with craft earned over years.
- **Red: the ten percent that is genuinely new.** The turn loop, context engineering, instruction layers, tool-result feedback, decision replay. Nothing in classic software maps to them.
- **Cyan: old craft at a new posting.** Model gateways, protocol adapters, event storage, and one kernel powering several faces — a CLI, an IDE, and a cloud surface. Built around the LLM, but the genuinely hard parts are rate limits, retries, failover, and version compatibility — nothing probabilistic.
- **Orange: the contamination belt.** Geographically inside traditional territory and doing old jobs, but its contracts were rewritten by the probability boundary — the pillar that actually bites, discussed separately below.

The orange pillar collects modules that used to be thoroughly familiar:

- The sandbox used to defend against outside attackers. Now it defends against your own agent — not malicious, just genuinely convinced that `rm -rf` is the correct next step.
- The process executor used to receive commands typed by humans. Now it receives strings generated by a model. The injection surface and the semantics of a hung process both change.
- Configuration used to be a formal language like YAML. AGENTS.md is natural language, and it fails in ways formal config never did.
- Logs and metrics used to locate errors. Now they must let you replay why the model thought a move was reasonable.

**Old modules. New contracts.** That is the contamination belt.

Root cause in one sentence:

> In a classic system the control flow is enumerable; an agent hands control of "what happens next" to a probabilistic function.

So the design target shifts wholesale: it used to be "the interface is correct"; now it is "**even when the model misunderstands, the system stays safe, recoverable, and observable**." Every design choice in the red and orange modules traces back to this sentence.

## How the series is built

The spine is the lifecycle of **one turn**: from the moment you press enter to the end of the task, every link taken apart. The two primary specimens are open source: codex (Rust, the complete industrial specimen) and opencode (TypeScript, a contrasting style). Each chapter follows the same pattern: one core question, the two implementations side by side, key code slices, and one diagram.

Planned chapters:

1. The life of one turn — what actually happens after you press enter
2. The core loop: plan, reflect, and a state machine that may not stop
3. The model boundary: protocols, rate limits, failover, cost
4. Context and memory: registers, RAM, and disk
5. Tools and execution: how the model actually acts
6. Sandbox and the trust boundary: defending against the insider
7. Instruction layers: programming in natural language
8. Extensions: what MCP, hooks, and plugins actually are
9. One kernel, many faces; and multi-agent
10. Testing a probabilistic system; closing principles for harness design

The order and the list will move around as the series is written.

For me the series is also an accounting exercise: after more than a decade of backend and platform work, which parts of agents are old skills wearing new clothes, and which are genuinely unseen problems. Taking that apart matters more than chasing terminology.

Next time: the life of one turn.
