from wasmtime import Engine, Linker, Module, Store, WasiConfig


def main():
    linker = Linker(Engine())
    linker.define_wasi()
    module = Module.from_file(linker.engine, "./builds/hello.wasm")

    config = WasiConfig()
    config.inherit_stdout()

    store = Store(linker.engine)
    store.set_wasi(config)
    instance = linker.instantiate(store, module)

    # _start is the default wasi main function
    start = instance.exports(store)["_start"]
    start(store)


if __name__ == '__main__':
    main()
