// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

import { themes } from 'prism-react-renderer';

const codeTheme = { light: themes.github, dark: themes.dracula };

/** @type {import('@docusaurus/types').Config} */
const config = {
    title: 'Codex',
    tagline: 'Instant, easy, and predictable shells and containers',
    url: 'https://www.khulnasoft',
    baseUrl: '/codex/docs/',
    onBrokenLinks: 'throw',
    onBrokenMarkdownLinks: 'warn',
    favicon: 'img/favicon.ico',
    trailingSlash: true,

    // GitHub pages deployment config.
    // If you aren't using GitHub pages, you don't need these.
    organizationName: 'khulnasoft', // Usually your GitHub org/user name.
    projectName: 'codex', // Usually your repo name.

    // Even if you don't use internalization, you can use this field to set useful
    // metadata like html lang. For example, if your site is Chinese, you may want
    // to replace "en" with "zh-Hans".
    markdown: {
        mermaid: true,
    },
    themes: [
        'docusaurus-theme-openapi-docs',
        '@docusaurus/theme-mermaid'
    ],
    i18n: {
        defaultLocale: 'en',
        locales: ['en'],
    },
    presets: [
        [
            'classic',
            /** @type {import('@docusaurus/preset-classic').Options} */
            ({
                docs: {
                    routeBasePath: '/',
                    sidebarPath: require.resolve('./sidebars.js'),
                    // Please change this to your repo.
                    // Remove this to remove the "edit this page" links.
                    docItemComponent: "@theme/ApiItem",
                    editUrl: "https://github.com/khulnasoft/codex/tree/main/docs/app/"
                },
                blog: false,
                theme: {
                    customCss: require.resolve('./src/css/custom.css'),
                },

                gtag: {
                    trackingID: 'G-PL4J94CXFK',
                    anonymizeIP: true,
                },
            } ),
        ],
    ],

    plugins: [
        [
            'docusaurus-plugin-openapi-docs',
            {
                id: 'api',
                docsPluginId: 'classic',
                config: {
                    nixhub: {
                        specPath: "specs/nixhub.yaml",
                        outputDir: "docs/nixhub",
                        
                    }
                }
            }
        ]
    ],

    themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
        ({
        navbar: {

            logo: {
                alt: 'Codex',
                src: 'img/codex_logo_light.svg',
                srcDark: 'img/codex_logo_dark.svg'
            },
            items: [
                {
                    to: 'https://cloud.khulnasoft',
                    label: 'Khulnasoft Cloud',
                    className: 'header-text-link',
                    position: 'left',
                  },
                {
                    href: 'https://discord.gg/khulnasoft',
                    // label: 'Discord',
                    className: 'header-discord-link',
                    position: 'right',
                },
                {
                    href: 'https://github.com/khulnasoft/codex',
                    // label: 'GitHub',
                    className: 'header-github-link',
                    position: 'right',
                },
            ],
        },
        footer: {
            links: [{
                    title: "Khulnasoft",
                    items: [{
                            label: "Khulnasoft",
                            href: "https://www.khulnasoft"
                        },
                        {
                            label: "Blog",
                            href: "https://www.khulnasoft/blog"
                        },
                        {
                            label: "Khulnasoft Cloud",
                            href: "https://cloud.khulnasoft"
                        }
                    ]
                },
                {
                    title: "Codex",
                    items: [{
                            label: "Home",
                            to: "https://www.khulnasoft/codex"
                        },
                        {
                            label: "Docs",
                            to: "https://www.khulnasoft/codex/docs/"
                        }
                    ]
                },

                {
                    title: "Community",
                    items: [

                        {
                            label: "Github",
                            href: "https://github.com/khulnasoft"
                        },
                        {
                            label: "Twitter",
                            href: "https://twitter.com/khulnasoft_com"
                        },
                        {
                            href: 'https://discord.gg/khulnasoft',
                            label: 'Discord',
                        },
                        {
                            href: "https://www.youtube.com/channel/UC7FwfJZbunZR2s-jG79vuTQ",
                            label: "Youtube"
                        }
                    ]
                }
            ],
            style: 'dark',
            copyright: `Copyright © ${new Date().getFullYear()} Khulnasoft, Inc.`,
        },
        colorMode: {
            respectPrefersColorScheme: true
        },
        algolia: {
            appId: 'J1RTMNIB0R',
            apiKey: 'b1bcbf465b384ccd6d986e85d6a62c28',
            indexName: 'khulnasoft',
            searchParameters: {},

        },
        prism: {
            theme: codeTheme.light,
            darkTheme: codeTheme.dark,
            additionalLanguages: ['bash', 'json'],
        },
    }),
};

export default config;
