// swift-tools-version: 6.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "LoveLang",
    platforms: [
        .macOS(.v15)
    ],
    dependencies: [

    ],
    targets: [
        .target(
            name: "LoveLangCore",
            dependencies: [

            ]),
        .testTarget(
            name: "LoveLangCoreTests",
            dependencies: [
                "LoveLangCore"
            ]),
        .executableTarget(
            name: "LoveLangCLI",
            dependencies: [
                "LoveLangCore"
            ]),

    ]
)
