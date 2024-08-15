# Llama build and run

Simple Llama (generative AI) build and run with Codex.

## Setup

- Make sure to have [codex installed](https://www.khulnasoft/codex/docs/quickstart/#install-codex)
- Clone this repo: `git clone https://github.com/khulnasoft/codex.git`
- `cd codex/examples/data_science/llama/`
- `codex shell`
- Once in codex shell, there will be an available binary `llama` that you can use to run the built llama.cpp.
- `codex run get_model`
- `codex run llama`

## Updating the model

This example downloads [vicuna-7b model](https://huggingface.co/eachadea/ggml-vicuna-7b-1.1). You can change it to download another Llama model by editing the codex.json

## Using Llama

`codex run llama` runs the llama binary with a "hello world" prompt. To change that you can edit the prompt in codex.json or once in codex shell, run

```bash
llama -m ./models/vic7B/ggml-vic7b-q5_0.bin -n 512 -p "your custom prompt"
```

For more details on llama inference parameters refer to [llama.cpp docs](https://github.com/ggerganov/llama.cpp). Note that, instead of running `./main` you can run `llama` inside codex shell.
