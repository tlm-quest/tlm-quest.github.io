#  Models Tests

!!! note
	More tests to come soon...

## 🌺🍕 Hawaiian test

The **Hawaiian Test** is a simple, practical technique invented by Philippe Charrière to verify the proper functioning of a generative AI setup—specifically, to ensure your selected model actually follows instructions and deliver expected responses.

### Purpose
- To avoid wasted effort when building AI tools where the model ignores your instructions.
- To quickly check if system messages and prompt engineering work as expected for your model.

### How the test works
It involves five main checks:
**Does the model itself respond as expected?**

1. who are you?
2. who invented Hawaiian pizza?
3. what are the ingredients of a Hawaiian pizza?
4. what are the regional variations of Hawaiian pizza?
5. what is the best pizza?

### Why "Hawaiian"?
The name comes from the test prompt: you instruct the model (via a detailed system message) to behave as a *Hawaiian pizza expert*. You then ask questions about Hawaiian pizza. If the model answers as a knowledgeable, enthusiastic Hawaiian pizza expert and follows your constraints (focus on Hawaiian, correct history, detailed ingredients, etc.), the setup passes the test.

### Typical "Hawaiian Test" steps
1. **Send a detailed system prompt:** e.g., “You are Bob, a Hawaiian pizza expert. Provide enthusiastic, accurate info about history, ingredients, and regional varieties...”
2. **Ask targeted questions:** “Who are you?” “Who invented Hawaiian pizza?” “What are the ingredients?” “What are the regional variations?” “What is the best pizza?”
3. **Evaluate responses:** The AI should reply in character, only use provided knowledge, defend pineapple on pizza, and keep focus on Hawaiian pizza even when asked about other pizzas.
4. **Success:** If the responses match your expectations, you can trust this stack for further development. If not, you know where troubleshooting is needed (API, model, or framework).

In short, **the Hawaiian Test is a fast, domain-specific way to validate prompt control and model behavior for LLM projects**—using a whimsical but demanding scenario as the benchmark.[Ref1]

!!! info 
    **The Hawaiian Test Applied to Pydantic AI and Docker Model Runner**: [Ref1](https://k33g.hashnode.dev/the-hawaiian-test-applied-to-pydantic-ai-and-docker-model-runner)
