package lib

import (
	"fmt"
	"testing"
)

var manifestV2 = `{
	"@context" : "http://iiif.io/api/presentation/2/context.json",
	"@id" : "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/manifest.json",
	"@type" : "sc:Manifest",
	"label" : "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade",
	"description" : "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade",
	"metadata" : [
	  {
		"label": "filename",
		"value": "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade"
	  }
	],
	"sequences" : [
	  {
		"@id" : "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/seq0.json",
		"@type" : "sc:Sequence",
		"label": "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade - sequence 1 0",
		"canvases": [
		  {
			"@id" : "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/can0.json",
			"@type" : "sc:Canvas",
			"label": "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade - image 0",
			"width" : 2905,
			"height" : 1920,
			"images": [
			  {
				"@id" : "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/anot0.json",
				"@type": "oa:Annotation",
				"motivation": "sc:painting",
				"on": "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/can0.json",
				"resource": {
				  "@id" : "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/full/full/0/native.jpg",
				  "@type" : "dctypes:Image",
				  "format" : "image/jpeg",
				  "width" : 2905,
				  "height" : 1920,
				  "service" : {
					"@context" : "http://iiif.io/api/image/2/context.json",
					"@id" : "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade",
					"profile" : "http://iiif.io/api/image/2/level1.json"
				  }
				}
			  }
			]
		  }
		]
	  }
	]
  }`

var manifestV3 = `{
    "@context": "http://iiif.io/api/presentation/3/context.json",
    "homepage": [
        {
            "format": "text/html",
            "id": "https://anet.be/record/uantwerpen/opacuantwerpen/c:lvd:727714",
            "label": {
                "en": [
                    "Catalogue record"
                ],
                "nl": [
                    "Catalogus record"
                ]
            },
            "type": "Text"
        }
    ],
    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/manifest",
    "items": [
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000001",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000001/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000001",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000001/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000001",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000001"
                ],
                "fr": [
                    "Image 00000001"
                ],
                "nl": [
                    "Image 00000001"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000001",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000002",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000002/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000002",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000002/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000002",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000002"
                ],
                "fr": [
                    "Image 00000002"
                ],
                "nl": [
                    "Image 00000002"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000002",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000003",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000003/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000003",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000003/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000003",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000003"
                ],
                "fr": [
                    "Image 00000003"
                ],
                "nl": [
                    "Image 00000003"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000003",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000004",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000004/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000004",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000004/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000004",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000004"
                ],
                "fr": [
                    "Image 00000004"
                ],
                "nl": [
                    "Image 00000004"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000004",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000005",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000005/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000005",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000005/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000005",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000005"
                ],
                "fr": [
                    "Image 00000005"
                ],
                "nl": [
                    "Image 00000005"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000005",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000006",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000006/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000006",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000006/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000006",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000006"
                ],
                "fr": [
                    "Image 00000006"
                ],
                "nl": [
                    "Image 00000006"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000006",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000007",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000007/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000007",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000007/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000007",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000007"
                ],
                "fr": [
                    "Image 00000007"
                ],
                "nl": [
                    "Image 00000007"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000007",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000008",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000008/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000008",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000008/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000008",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000008"
                ],
                "fr": [
                    "Image 00000008"
                ],
                "nl": [
                    "Image 00000008"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000008",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000009",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000009/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000009",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000009/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000009",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000009"
                ],
                "fr": [
                    "Image 00000009"
                ],
                "nl": [
                    "Image 00000009"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000009",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000010",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000010/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000010",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000010/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000010",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000010"
                ],
                "fr": [
                    "Image 00000010"
                ],
                "nl": [
                    "Image 00000010"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000010",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000011",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000011/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000011",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000011/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000011",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000011"
                ],
                "fr": [
                    "Image 00000011"
                ],
                "nl": [
                    "Image 00000011"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000011",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000012",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000012/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000012",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000012/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000012",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000012"
                ],
                "fr": [
                    "Image 00000012"
                ],
                "nl": [
                    "Image 00000012"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000012",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000013",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000013/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000013",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000013/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000013",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000013"
                ],
                "fr": [
                    "Image 00000013"
                ],
                "nl": [
                    "Image 00000013"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000013",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000014",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000014/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000014",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000014/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000014",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000014"
                ],
                "fr": [
                    "Image 00000014"
                ],
                "nl": [
                    "Image 00000014"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000014",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000015",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000015/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000015",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000015/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000015",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000015"
                ],
                "fr": [
                    "Image 00000015"
                ],
                "nl": [
                    "Image 00000015"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000015",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000016",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000016/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000016",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000016/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000016",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000016"
                ],
                "fr": [
                    "Image 00000016"
                ],
                "nl": [
                    "Image 00000016"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000016",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000017",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000017/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000017",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000017/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000017",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000017"
                ],
                "fr": [
                    "Image 00000017"
                ],
                "nl": [
                    "Image 00000017"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000017",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000018",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000018/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000018",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000018/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000018",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000018"
                ],
                "fr": [
                    "Image 00000018"
                ],
                "nl": [
                    "Image 00000018"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000018",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000019",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000019/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000019",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000019/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000019",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000019"
                ],
                "fr": [
                    "Image 00000019"
                ],
                "nl": [
                    "Image 00000019"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000019",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000020",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000020/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000020",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000020/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000020",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000020"
                ],
                "fr": [
                    "Image 00000020"
                ],
                "nl": [
                    "Image 00000020"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000020",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000021",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000021/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000021",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000021/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000021",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000021"
                ],
                "fr": [
                    "Image 00000021"
                ],
                "nl": [
                    "Image 00000021"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000021",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000022",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000022/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000022",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000022/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000022",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000022"
                ],
                "fr": [
                    "Image 00000022"
                ],
                "nl": [
                    "Image 00000022"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000022",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000023",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000023/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000023",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000023/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000023",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000023"
                ],
                "fr": [
                    "Image 00000023"
                ],
                "nl": [
                    "Image 00000023"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000023",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000024",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000024/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000024",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000024/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000024",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000024"
                ],
                "fr": [
                    "Image 00000024"
                ],
                "nl": [
                    "Image 00000024"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000024",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000025",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000025/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000025",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000025/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000025",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000025"
                ],
                "fr": [
                    "Image 00000025"
                ],
                "nl": [
                    "Image 00000025"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000025",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000026",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000026/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000026",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000026/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000026",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000026"
                ],
                "fr": [
                    "Image 00000026"
                ],
                "nl": [
                    "Image 00000026"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000026",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000027",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000027/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000027",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000027/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000027",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000027"
                ],
                "fr": [
                    "Image 00000027"
                ],
                "nl": [
                    "Image 00000027"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000027",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000028",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000028/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000028",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000028/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000028",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000028"
                ],
                "fr": [
                    "Image 00000028"
                ],
                "nl": [
                    "Image 00000028"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000028",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000029",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000029/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000029",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000029/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000029",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000029"
                ],
                "fr": [
                    "Image 00000029"
                ],
                "nl": [
                    "Image 00000029"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000029",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000030",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000030/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000030",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000030/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000030",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000030"
                ],
                "fr": [
                    "Image 00000030"
                ],
                "nl": [
                    "Image 00000030"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000030",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000031",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000031/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000031",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000031/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000031",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000031"
                ],
                "fr": [
                    "Image 00000031"
                ],
                "nl": [
                    "Image 00000031"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000031",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000032",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000032/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000032",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000032/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000032",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000032"
                ],
                "fr": [
                    "Image 00000032"
                ],
                "nl": [
                    "Image 00000032"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000032",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000033",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000033/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000033",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000033/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000033",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000033"
                ],
                "fr": [
                    "Image 00000033"
                ],
                "nl": [
                    "Image 00000033"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000033",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000034",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000034/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000034",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000034/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000034",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000034"
                ],
                "fr": [
                    "Image 00000034"
                ],
                "nl": [
                    "Image 00000034"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000034",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000035",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000035/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000035",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000035/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000035",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000035"
                ],
                "fr": [
                    "Image 00000035"
                ],
                "nl": [
                    "Image 00000035"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000035",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000036",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000036/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000036",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000036/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000036",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000036"
                ],
                "fr": [
                    "Image 00000036"
                ],
                "nl": [
                    "Image 00000036"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000036",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000037",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000037/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000037",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000037/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000037",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000037"
                ],
                "fr": [
                    "Image 00000037"
                ],
                "nl": [
                    "Image 00000037"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000037",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000038",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000038/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000038",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000038/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000038",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000038"
                ],
                "fr": [
                    "Image 00000038"
                ],
                "nl": [
                    "Image 00000038"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000038",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000039",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000039/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000039",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000039/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000039",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000039"
                ],
                "fr": [
                    "Image 00000039"
                ],
                "nl": [
                    "Image 00000039"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000039",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000040",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000040/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000040",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000040/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000040",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000040"
                ],
                "fr": [
                    "Image 00000040"
                ],
                "nl": [
                    "Image 00000040"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000040",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000041",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000041/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000041",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000041/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000041",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000041"
                ],
                "fr": [
                    "Image 00000041"
                ],
                "nl": [
                    "Image 00000041"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000041",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000042",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000042/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000042",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000042/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000042",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000042"
                ],
                "fr": [
                    "Image 00000042"
                ],
                "nl": [
                    "Image 00000042"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000042",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000043",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000043/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000043",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000043/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000043",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000043"
                ],
                "fr": [
                    "Image 00000043"
                ],
                "nl": [
                    "Image 00000043"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000043",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000044",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000044/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000044",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000044/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000044",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000044"
                ],
                "fr": [
                    "Image 00000044"
                ],
                "nl": [
                    "Image 00000044"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000044",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000045",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000045/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000045",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000045/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000045",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000045"
                ],
                "fr": [
                    "Image 00000045"
                ],
                "nl": [
                    "Image 00000045"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000045",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000046",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000046/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000046",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000046/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000046",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000046"
                ],
                "fr": [
                    "Image 00000046"
                ],
                "nl": [
                    "Image 00000046"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000046",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000047",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000047/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000047",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000047/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000047",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000047"
                ],
                "fr": [
                    "Image 00000047"
                ],
                "nl": [
                    "Image 00000047"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000047",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        },
        {
            "height": 600,
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000048",
            "items": [
                {
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000048/1",
                    "items": [
                        {
                            "body": {
                                "format": "image/jpeg",
                                "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000048",
                                "type": "Image"
                            },
                            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000048/1/image",
                            "motivation": "painting",
                            "target": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvasbase/00000048",
                            "type": "Annotation"
                        }
                    ],
                    "type": "AnnotationPage"
                }
            ],
            "label": {
                "en": [
                    "Image 00000048"
                ],
                "fr": [
                    "Image 00000048"
                ],
                "nl": [
                    "Image 00000048"
                ]
            },
            "thumbnail": [
                {
                    "format": "image/jpg",
                    "height": 600,
                    "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000048",
                    "type": "Image",
                    "width": 400
                }
            ],
            "type": "Canvas",
            "width": 400
        }
    ],
    "label": {
        "en": [
            "De veinzende Torquatus, treurspel"
        ],
        "nl": [
            "De veinzende Torquatus, treurspel"
        ]
    },
    "logo": [
        {
            "format": "image/svg+xml",
            "height": 50,
            "id": "https://anet.be/desktop/uantwerpen/static/Banner_website_UAntwerpen_Bibliotheek_01.svg",
            "type": "Image",
            "width": 100
        }
    ],
    "metadata": [
        {
            "label": {
                "en": [
                    "Title"
                ],
                "nl": [
                    "Titel"
                ]
            },
            "value": {
                "en": [
                    "De veinzende Torquatus, treurspel"
                ],
                "nl": [
                    "De veinzende Torquatus, treurspel"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Creator"
                ],
                "nl": [
                    "Auteur"
                ]
            },
            "value": {
                "en": [
                    " Brandt, Geeraert"
                ],
                "nl": [
                    " Brandt, Geeraert"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Impressum"
                ],
                "nl": [
                    "Impressum"
                ]
            },
            "value": {
                "en": [
                    "Amsteldam : gedruckt by Jacob Lescaille, 1645"
                ],
                "nl": [
                    "Amsteldam : gedruckt by Jacob Lescaille, 1645"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Edition"
                ],
                "nl": [
                    "Uitgave"
                ]
            },
            "value": {
                "en": [
                    ""
                ],
                "nl": [
                    ""
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Collation"
                ],
                "nl": [
                    "Collatie"
                ]
            },
            "value": {
                "en": [
                    "[10], 71, [5] p. : front. ; quarto ; π1 \u003csup\u003e2\u003c/sup\u003eπ\u003csup\u003e4\u003c/sup\u003e A-I\u003csup\u003e4\u003c/sup\u003e a\u003csup\u003e2\u003c/sup\u003e"
                ],
                "nl": [
                    "[10], 71, [5] p. : front. ; quarto ; π1 \u003csup\u003e2\u003c/sup\u003eπ\u003csup\u003e4\u003c/sup\u003e A-I\u003csup\u003e4\u003c/sup\u003e a\u003csup\u003e2\u003c/sup\u003e"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Library"
                ],
                "nl": [
                    "Bibliotheek"
                ]
            },
            "value": {
                "en": [
                    "UAntwerpen - Bibliotheek Stadscampus"
                ],
                "nl": [
                    "UAntwerpen - Bibliotheek Stadscampus"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Call number"
                ],
                "nl": [
                    "Plaatskenmerk"
                ]
            },
            "value": {
                "en": [
                    "MAG-P 12.2212"
                ],
                "nl": [
                    "MAG-P 12.2212"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Barcode"
                ],
                "nl": [
                    "Streepjescode"
                ]
            },
            "value": {
                "en": [
                    "03020651880"
                ],
                "nl": [
                    "03020651880"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Collection"
                ],
                "nl": [
                    "Collectie"
                ]
            },
            "value": {
                "en": [
                    "Collectie preciosa (digitaal)"
                ],
                "nl": [
                    "Collectie preciosa (digitaal)"
                ]
            }
        },
        {
            "label": {
                "en": [
                    "Collection"
                ],
                "nl": [
                    "Collectie"
                ]
            },
            "value": {
                "en": [
                    "Collectie Centrum Renaissancedrama, UAntwerpen"
                ],
                "nl": [
                    "Collectie Centrum Renaissancedrama, UAntwerpen"
                ]
            }
        }
    ],
    "provider": [
        {
            "homepage": [
                {
                    "format": "text/html",
                    "id": "https://uantwerpen.be/library",
                    "label": {
                        "en": [
                            "Homepage - University of Antwerp Library"
                        ]
                    },
                    "type": "Text"
                }
            ],
            "id": "https://uantwerpen.be/ ",
            "label": {
                "en": [
                    "University of Antwerp Library"
                ]
            },
            "logo": [
                {
                    "format": "image/svg+xml",
                    "height": 50,
                    "id": "https://anet.be/desktop/uantwerpen/static/Banner_website_UAntwerpen_Bibliotheek_01.svg",
                    "type": "Image",
                    "width": 50
                }
            ],
            "type": "Agent"
        }
    ],
    "rendering": [
        {
            "format": "application/vnd.oasis.opendocument.presentation",
            "id": "https://anet.be/docman/uapreciosa/91396f/1.odp",
            "label": {
                "en": [
                    "Download as ODP file"
                ],
                "fr": [
                    "Télécharger en fichier ODP"
                ],
                "nl": [
                    "Download als ODP bestand"
                ]
            },
            "type": "Dataset"
        },
        {
            "format": "application/pdf",
            "id": "https://anet.be/docman/uapreciosa/23ee73/1.pdf",
            "label": {
                "en": [
                    "Download as PDF file"
                ],
                "fr": [
                    "Télécharger en fichier PDF"
                ],
                "nl": [
                    "Download als PDF bestand"
                ]
            },
            "type": "Dataset"
        },
        {
            "format": "application/zip",
            "id": "https://anet.be/docman/uapreciosa/2a4191/1.zip",
            "label": {
                "en": [
                    "Download as ZIP file"
                ],
                "fr": [
                    "Télécharger en fichier ZIP"
                ],
                "nl": [
                    "Download als ZIP bestand"
                ]
            },
            "type": "Dataset"
        },
        {
            "id": "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/sqlite",
            "label": {
                "en": [
                    "SQLite"
                ],
                "fr": [
                    "SQLite"
                ],
                "nl": [
                    "SQLite"
                ]
            },
            "type": "Dataset"
        }
    ],
    "requiredStatement": {
        "label": {
            "en": [
                "Attribution"
            ],
            "nl": [
                "Attributie"
            ]
        },
        "value": {
            "en": [
                "Provided by University of Antwerp Libraries"
            ],
            "nl": [
                "Beschikbaar gesteld door de bibliotheek van UAntwerpen"
            ]
        }
    },
    "rights": "http://creativecommons.org/licenses/by/4.0/",
    "seeAlso": [
        {
            "format": "text/xml",
            "id": "https://anet.be/oai/catunicat/server.phtml?verb=GetRecord\u0026metadataPrefix=marc21\u0026identifier=c:lvd:727714",
            "label": {
                "en": [
                    "Bibliographic Description in MARCXML"
                ],
                "nl": [
                    "Bibliographische beschrijving in MARCXML"
                ]
            },
            "type": "Dataset"
        },
        {
            "format": "text/xml",
            "id": "https://anet.be/oai/catgeneric/server.phtml?verb=GetRecord\u0026metadataPrefix=catxml\u0026identifier=c:lvd:727714",
            "label": {
                "en": [
                    "Bibliographic Description in CATXML"
                ],
                "nl": [
                    "Bibliographische beschrijving in CATXML"
                ]
            },
            "type": "Dataset"
        }
    ],
    "summary": {
        "en": [
            "De veinzende Torquatus, treurspel / Brandt, Geeraert. - Amsteldam : gedruckt by Jacob Lescaille, 1645. - [10], 71, [5] p. : front. ; quarto ; π1 \u003csup\u003e2\u003c/sup\u003eπ\u003csup\u003e4\u003c/sup\u003e A-I\u003csup\u003e4\u003c/sup\u003e a\u003csup\u003e2\u003c/sup\u003e. -  Bibliographic reference: STCN 842313370"
        ],
        "nl": [
            "De veinzende Torquatus, treurspel / Brandt, Geeraert. - Amsteldam : gedruckt by Jacob Lescaille, 1645. - [10], 71, [5] p. : front. ; quarto ; π1 \u003csup\u003e2\u003c/sup\u003eπ\u003csup\u003e4\u003c/sup\u003e A-I\u003csup\u003e4\u003c/sup\u003e a\u003csup\u003e2\u003c/sup\u003e. -  Bibliografische referentie: STCN 842313370"
        ]
    },
    "type": "Manifest"
}`

func TestManifestV3(t *testing.T) {
	result, err := ParseManifest([]byte(manifestV3))
	if err != nil {
		t.Errorf(fmt.Sprintf("Manifest3 error[%s]\n", err))
	}
	summaryen := "De veinzende Torquatus, treurspel / Brandt, Geeraert. - Amsteldam : gedruckt by Jacob Lescaille, 1645. - [10], 71, [5] p. : front. ; quarto ; π1 <sup>2</sup>π<sup>4</sup> A-I<sup>4</sup> a<sup>2</sup>. -  Bibliographic reference: STCN 842313370"
	expected := ManifestInfo{
		Label:      "De veinzende Torquatus, treurspel",
		Summary:    summaryen,
		Thumbnail:  "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/canvas/00000001",
		Identifier: "https://anet.be/iiif/cc09b2bc22e874e9165253f02295f8b8a0420ee0/manifest"}
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n%v\nExpected: \n%v\n", result, expected))
	}
}

func TestManifestV2(t *testing.T) {
	result, err := ParseManifest([]byte(manifestV2))
	if err != nil {
		t.Errorf(fmt.Sprintf("Manifest2 error[%s]\n", err))
	}
	expected := ManifestInfo{
		Label:      "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade",
		Summary:    "iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade",
		Thumbnail:  "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/full/full/0/native.jpg",
		Identifier: "http://nationalmuseumse.iiifhosting.com/iiif/83542114f728ce16ea0054cdac4f13bb3702bb2de1e3972055412920f4eb1ade/manifest.json"}
	if result != expected {
		t.Errorf(fmt.Sprintf("\nResult: \n%v\nExpected: \n%v\n", result, expected))
	}
}
