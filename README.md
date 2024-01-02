# Angle Test

Represent an angle + a quantized vector length as a single byte.

Current function:

```go
func angleLength(input uint8) (uint8, uint8) {
    const twopi = 2 * math.Pi
    const steps = 10
    if input == 0 {
        return 0, 0
    }
    length := (input / steps) * steps
    inputFloat64 := float64(input) / 255.0
    angleByte := uint8(math.Mod(inputFloat64*steps*twopi, twopi) * (255.0 / twopi))
    return angleByte, length
}
```

Input and output values + images:

![input 0](img/vector000.png)

    input 0 -> angle 0.000000, length 0.000000

![input 1](img/vector001.png)

    input 1 -> angle 0.039216, length 0.000000

![input 2](img/vector002.png)

    input 2 -> angle 0.078431, length 0.000000

![input 3](img/vector003.png)

    input 3 -> angle 0.117647, length 0.000000

![input 4](img/vector004.png)

    input 4 -> angle 0.156863, length 0.000000

![input 5](img/vector005.png)

    input 5 -> angle 0.196078, length 0.000000

![input 6](img/vector006.png)

    input 6 -> angle 0.235294, length 0.000000

![input 7](img/vector007.png)

    input 7 -> angle 0.274510, length 0.000000

![input 8](img/vector008.png)

    input 8 -> angle 0.313725, length 0.000000

![input 9](img/vector009.png)

    input 9 -> angle 0.349020, length 0.000000

![input 10](img/vector010.png)

    input 10 -> angle 0.392157, length 0.039216

![input 11](img/vector011.png)

    input 11 -> angle 0.431373, length 0.039216

![input 12](img/vector012.png)

    input 12 -> angle 0.470588, length 0.039216

![input 13](img/vector013.png)

    input 13 -> angle 0.505882, length 0.039216

![input 14](img/vector014.png)

    input 14 -> angle 0.549020, length 0.039216

![input 15](img/vector015.png)

    input 15 -> angle 0.588235, length 0.039216

![input 16](img/vector016.png)

    input 16 -> angle 0.627451, length 0.039216

![input 17](img/vector017.png)

    input 17 -> angle 0.662745, length 0.039216

![input 18](img/vector018.png)

    input 18 -> angle 0.701961, length 0.039216

![input 19](img/vector019.png)

    input 19 -> angle 0.745098, length 0.039216

![input 20](img/vector020.png)

    input 20 -> angle 0.784314, length 0.078431

![input 21](img/vector021.png)

    input 21 -> angle 0.819608, length 0.078431

![input 22](img/vector022.png)

    input 22 -> angle 0.862745, length 0.078431

![input 23](img/vector023.png)

    input 23 -> angle 0.898039, length 0.078431

![input 24](img/vector024.png)

    input 24 -> angle 0.941176, length 0.078431

![input 25](img/vector025.png)

    input 25 -> angle 0.976471, length 0.078431

![input 26](img/vector026.png)

    input 26 -> angle 0.015686, length 0.078431

![input 27](img/vector027.png)

    input 27 -> angle 0.054902, length 0.078431

![input 28](img/vector028.png)

    input 28 -> angle 0.098039, length 0.078431

![input 29](img/vector029.png)

    input 29 -> angle 0.133333, length 0.078431

![input 30](img/vector030.png)

    input 30 -> angle 0.176471, length 0.117647

![input 31](img/vector031.png)

    input 31 -> angle 0.211765, length 0.117647

![input 32](img/vector032.png)

    input 32 -> angle 0.250980, length 0.117647

![input 33](img/vector033.png)

    input 33 -> angle 0.294118, length 0.117647

![input 34](img/vector034.png)

    input 34 -> angle 0.329412, length 0.117647

![input 35](img/vector035.png)

    input 35 -> angle 0.372549, length 0.117647

![input 36](img/vector036.png)

    input 36 -> angle 0.407843, length 0.117647

![input 37](img/vector037.png)

    input 37 -> angle 0.450980, length 0.117647

![input 38](img/vector038.png)

    input 38 -> angle 0.490196, length 0.117647

![input 39](img/vector039.png)

    input 39 -> angle 0.529412, length 0.117647

![input 40](img/vector040.png)

    input 40 -> angle 0.568627, length 0.156863

![input 41](img/vector041.png)

    input 41 -> angle 0.603922, length 0.156863

![input 42](img/vector042.png)

    input 42 -> angle 0.643137, length 0.156863

![input 43](img/vector043.png)

    input 43 -> angle 0.686275, length 0.156863

![input 44](img/vector044.png)

    input 44 -> angle 0.725490, length 0.156863

![input 45](img/vector045.png)

    input 45 -> angle 0.764706, length 0.156863

![input 46](img/vector046.png)

    input 46 -> angle 0.800000, length 0.156863

![input 47](img/vector047.png)

    input 47 -> angle 0.843137, length 0.156863

![input 48](img/vector048.png)

    input 48 -> angle 0.882353, length 0.156863

![input 49](img/vector049.png)

    input 49 -> angle 0.917647, length 0.156863

![input 50](img/vector050.png)

    input 50 -> angle 0.956863, length 0.196078

![input 51](img/vector051.png)

    input 51 -> angle 0.000000, length 0.196078

![input 52](img/vector052.png)

    input 52 -> angle 0.035294, length 0.196078

![input 53](img/vector053.png)

    input 53 -> angle 0.078431, length 0.196078

![input 54](img/vector054.png)

    input 54 -> angle 0.113725, length 0.196078

![input 55](img/vector055.png)

    input 55 -> angle 0.156863, length 0.196078

![input 56](img/vector056.png)

    input 56 -> angle 0.196078, length 0.196078

![input 57](img/vector057.png)

    input 57 -> angle 0.235294, length 0.196078

![input 58](img/vector058.png)

    input 58 -> angle 0.270588, length 0.196078

![input 59](img/vector059.png)

    input 59 -> angle 0.313725, length 0.196078

![input 60](img/vector060.png)

    input 60 -> angle 0.352941, length 0.235294

![input 61](img/vector061.png)

    input 61 -> angle 0.392157, length 0.235294

![input 62](img/vector062.png)

    input 62 -> angle 0.427451, length 0.235294

![input 63](img/vector063.png)

    input 63 -> angle 0.470588, length 0.235294

![input 64](img/vector064.png)

    input 64 -> angle 0.505882, length 0.235294

![input 65](img/vector065.png)

    input 65 -> angle 0.545098, length 0.235294

![input 66](img/vector066.png)

    input 66 -> angle 0.588235, length 0.235294

![input 67](img/vector067.png)

    input 67 -> angle 0.623529, length 0.235294

![input 68](img/vector068.png)

    input 68 -> angle 0.662745, length 0.235294

![input 69](img/vector069.png)

    input 69 -> angle 0.701961, length 0.235294

![input 70](img/vector070.png)

    input 70 -> angle 0.745098, length 0.274510

![input 71](img/vector071.png)

    input 71 -> angle 0.784314, length 0.274510

![input 72](img/vector072.png)

    input 72 -> angle 0.819608, length 0.274510

![input 73](img/vector073.png)

    input 73 -> angle 0.858824, length 0.274510

![input 74](img/vector074.png)

    input 74 -> angle 0.901961, length 0.274510

![input 75](img/vector075.png)

    input 75 -> angle 0.941176, length 0.274510

![input 76](img/vector076.png)

    input 76 -> angle 0.980392, length 0.274510

![input 77](img/vector077.png)

    input 77 -> angle 0.015686, length 0.274510

![input 78](img/vector078.png)

    input 78 -> angle 0.058824, length 0.274510

![input 79](img/vector079.png)

    input 79 -> angle 0.098039, length 0.274510

![input 80](img/vector080.png)

    input 80 -> angle 0.137255, length 0.313725

![input 81](img/vector081.png)

    input 81 -> angle 0.176471, length 0.313725

![input 82](img/vector082.png)

    input 82 -> angle 0.211765, length 0.313725

![input 83](img/vector083.png)

    input 83 -> angle 0.250980, length 0.313725

![input 84](img/vector084.png)

    input 84 -> angle 0.290196, length 0.313725

![input 85](img/vector085.png)

    input 85 -> angle 0.329412, length 0.313725

![input 86](img/vector086.png)

    input 86 -> angle 0.372549, length 0.313725

![input 87](img/vector087.png)

    input 87 -> angle 0.411765, length 0.313725

![input 88](img/vector088.png)

    input 88 -> angle 0.450980, length 0.313725

![input 89](img/vector089.png)

    input 89 -> angle 0.490196, length 0.313725

![input 90](img/vector090.png)

    input 90 -> angle 0.529412, length 0.352941

![input 91](img/vector091.png)

    input 91 -> angle 0.568627, length 0.352941

![input 92](img/vector092.png)

    input 92 -> angle 0.603922, length 0.352941

![input 93](img/vector093.png)

    input 93 -> angle 0.643137, length 0.352941

![input 94](img/vector094.png)

    input 94 -> angle 0.686275, length 0.352941

![input 95](img/vector095.png)

    input 95 -> angle 0.725490, length 0.352941

![input 96](img/vector096.png)

    input 96 -> angle 0.764706, length 0.352941

![input 97](img/vector097.png)

    input 97 -> angle 0.800000, length 0.352941

![input 98](img/vector098.png)

    input 98 -> angle 0.839216, length 0.352941

![input 99](img/vector099.png)

    input 99 -> angle 0.878431, length 0.352941

![input 100](img/vector100.png)

    input 100 -> angle 0.917647, length 0.392157

![input 101](img/vector101.png)

    input 101 -> angle 0.956863, length 0.392157

![input 102](img/vector102.png)

    input 102 -> angle 0.000000, length 0.392157

![input 103](img/vector103.png)

    input 103 -> angle 0.035294, length 0.392157

![input 104](img/vector104.png)

    input 104 -> angle 0.074510, length 0.392157

![input 105](img/vector105.png)

    input 105 -> angle 0.113725, length 0.392157

![input 106](img/vector106.png)

    input 106 -> angle 0.156863, length 0.392157

![input 107](img/vector107.png)

    input 107 -> angle 0.196078, length 0.392157

![input 108](img/vector108.png)

    input 108 -> angle 0.231373, length 0.392157

![input 109](img/vector109.png)

    input 109 -> angle 0.270588, length 0.392157

![input 110](img/vector110.png)

    input 110 -> angle 0.313725, length 0.431373

![input 111](img/vector111.png)

    input 111 -> angle 0.352941, length 0.431373

![input 112](img/vector112.png)

    input 112 -> angle 0.392157, length 0.431373

![input 113](img/vector113.png)

    input 113 -> angle 0.431373, length 0.431373

![input 114](img/vector114.png)

    input 114 -> angle 0.470588, length 0.431373

![input 115](img/vector115.png)

    input 115 -> angle 0.505882, length 0.431373

![input 116](img/vector116.png)

    input 116 -> angle 0.545098, length 0.431373

![input 117](img/vector117.png)

    input 117 -> angle 0.584314, length 0.431373

![input 118](img/vector118.png)

    input 118 -> angle 0.627451, length 0.431373

![input 119](img/vector119.png)

    input 119 -> angle 0.666667, length 0.431373

![input 120](img/vector120.png)

    input 120 -> angle 0.705882, length 0.470588

![input 121](img/vector121.png)

    input 121 -> angle 0.745098, length 0.470588

![input 122](img/vector122.png)

    input 122 -> angle 0.784314, length 0.470588

![input 123](img/vector123.png)

    input 123 -> angle 0.819608, length 0.470588

![input 124](img/vector124.png)

    input 124 -> angle 0.858824, length 0.470588

![input 125](img/vector125.png)

    input 125 -> angle 0.898039, length 0.470588

![input 126](img/vector126.png)

    input 126 -> angle 0.941176, length 0.470588

![input 127](img/vector127.png)

    input 127 -> angle 0.980392, length 0.470588

![input 128](img/vector128.png)

    input 128 -> angle 0.015686, length 0.470588

![input 129](img/vector129.png)

    input 129 -> angle 0.054902, length 0.470588

![input 130](img/vector130.png)

    input 130 -> angle 0.094118, length 0.509804

![input 131](img/vector131.png)

    input 131 -> angle 0.133333, length 0.509804

![input 132](img/vector132.png)

    input 132 -> angle 0.176471, length 0.509804

![input 133](img/vector133.png)

    input 133 -> angle 0.211765, length 0.509804

![input 134](img/vector134.png)

    input 134 -> angle 0.250980, length 0.509804

![input 135](img/vector135.png)

    input 135 -> angle 0.290196, length 0.509804

![input 136](img/vector136.png)

    input 136 -> angle 0.329412, length 0.509804

![input 137](img/vector137.png)

    input 137 -> angle 0.368627, length 0.509804

![input 138](img/vector138.png)

    input 138 -> angle 0.407843, length 0.509804

![input 139](img/vector139.png)

    input 139 -> angle 0.447059, length 0.509804

![input 140](img/vector140.png)

    input 140 -> angle 0.490196, length 0.549020

![input 141](img/vector141.png)

    input 141 -> angle 0.529412, length 0.549020

![input 142](img/vector142.png)

    input 142 -> angle 0.568627, length 0.549020

![input 143](img/vector143.png)

    input 143 -> angle 0.603922, length 0.549020

![input 144](img/vector144.png)

    input 144 -> angle 0.643137, length 0.549020

![input 145](img/vector145.png)

    input 145 -> angle 0.682353, length 0.549020

![input 146](img/vector146.png)

    input 146 -> angle 0.721569, length 0.549020

![input 147](img/vector147.png)

    input 147 -> angle 0.760784, length 0.549020

![input 148](img/vector148.png)

    input 148 -> angle 0.803922, length 0.549020

![input 149](img/vector149.png)

    input 149 -> angle 0.843137, length 0.549020

![input 150](img/vector150.png)

    input 150 -> angle 0.882353, length 0.588235

![input 151](img/vector151.png)

    input 151 -> angle 0.921569, length 0.588235

![input 152](img/vector152.png)

    input 152 -> angle 0.960784, length 0.588235

![input 153](img/vector153.png)

    input 153 -> angle 0.000000, length 0.588235

![input 154](img/vector154.png)

    input 154 -> angle 0.035294, length 0.588235

![input 155](img/vector155.png)

    input 155 -> angle 0.074510, length 0.588235

![input 156](img/vector156.png)

    input 156 -> angle 0.117647, length 0.588235

![input 157](img/vector157.png)

    input 157 -> angle 0.156863, length 0.588235

![input 158](img/vector158.png)

    input 158 -> angle 0.196078, length 0.588235

![input 159](img/vector159.png)

    input 159 -> angle 0.235294, length 0.588235

![input 160](img/vector160.png)

    input 160 -> angle 0.274510, length 0.627451

![input 161](img/vector161.png)

    input 161 -> angle 0.313725, length 0.627451

![input 162](img/vector162.png)

    input 162 -> angle 0.352941, length 0.627451

![input 163](img/vector163.png)

    input 163 -> angle 0.392157, length 0.627451

![input 164](img/vector164.png)

    input 164 -> angle 0.427451, length 0.627451

![input 165](img/vector165.png)

    input 165 -> angle 0.466667, length 0.627451

![input 166](img/vector166.png)

    input 166 -> angle 0.505882, length 0.627451

![input 167](img/vector167.png)

    input 167 -> angle 0.545098, length 0.627451

![input 168](img/vector168.png)

    input 168 -> angle 0.584314, length 0.627451

![input 169](img/vector169.png)

    input 169 -> angle 0.623529, length 0.627451

![input 170](img/vector170.png)

    input 170 -> angle 0.662745, length 0.666667

![input 171](img/vector171.png)

    input 171 -> angle 0.701961, length 0.666667

![input 172](img/vector172.png)

    input 172 -> angle 0.745098, length 0.666667

![input 173](img/vector173.png)

    input 173 -> angle 0.784314, length 0.666667

![input 174](img/vector174.png)

    input 174 -> angle 0.823529, length 0.666667

![input 175](img/vector175.png)

    input 175 -> angle 0.862745, length 0.666667

![input 176](img/vector176.png)

    input 176 -> angle 0.901961, length 0.666667

![input 177](img/vector177.png)

    input 177 -> angle 0.941176, length 0.666667

![input 178](img/vector178.png)

    input 178 -> angle 0.980392, length 0.666667

![input 179](img/vector179.png)

    input 179 -> angle 0.015686, length 0.666667

![input 180](img/vector180.png)

    input 180 -> angle 0.058824, length 0.705882

![input 181](img/vector181.png)

    input 181 -> angle 0.098039, length 0.705882

![input 182](img/vector182.png)

    input 182 -> angle 0.137255, length 0.705882

![input 183](img/vector183.png)

    input 183 -> angle 0.176471, length 0.705882

![input 184](img/vector184.png)

    input 184 -> angle 0.211765, length 0.705882

![input 185](img/vector185.png)

    input 185 -> angle 0.250980, length 0.705882

![input 186](img/vector186.png)

    input 186 -> angle 0.290196, length 0.705882

![input 187](img/vector187.png)

    input 187 -> angle 0.329412, length 0.705882

![input 188](img/vector188.png)

    input 188 -> angle 0.372549, length 0.705882

![input 189](img/vector189.png)

    input 189 -> angle 0.411765, length 0.705882

![input 190](img/vector190.png)

    input 190 -> angle 0.450980, length 0.745098

![input 191](img/vector191.png)

    input 191 -> angle 0.490196, length 0.745098

![input 192](img/vector192.png)

    input 192 -> angle 0.529412, length 0.745098

![input 193](img/vector193.png)

    input 193 -> angle 0.568627, length 0.745098

![input 194](img/vector194.png)

    input 194 -> angle 0.603922, length 0.745098

![input 195](img/vector195.png)

    input 195 -> angle 0.643137, length 0.745098

![input 196](img/vector196.png)

    input 196 -> angle 0.682353, length 0.745098

![input 197](img/vector197.png)

    input 197 -> angle 0.721569, length 0.745098

![input 198](img/vector198.png)

    input 198 -> angle 0.760784, length 0.745098

![input 199](img/vector199.png)

    input 199 -> angle 0.800000, length 0.745098

![input 200](img/vector200.png)

    input 200 -> angle 0.839216, length 0.784314

![input 201](img/vector201.png)

    input 201 -> angle 0.878431, length 0.784314

![input 202](img/vector202.png)

    input 202 -> angle 0.917647, length 0.784314

![input 203](img/vector203.png)

    input 203 -> angle 0.956863, length 0.784314

![input 204](img/vector204.png)

    input 204 -> angle 0.000000, length 0.784314

![input 205](img/vector205.png)

    input 205 -> angle 0.035294, length 0.784314

![input 206](img/vector206.png)

    input 206 -> angle 0.074510, length 0.784314

![input 207](img/vector207.png)

    input 207 -> angle 0.113725, length 0.784314

![input 208](img/vector208.png)

    input 208 -> angle 0.152941, length 0.784314

![input 209](img/vector209.png)

    input 209 -> angle 0.192157, length 0.784314

![input 210](img/vector210.png)

    input 210 -> angle 0.231373, length 0.823529

![input 211](img/vector211.png)

    input 211 -> angle 0.270588, length 0.823529

![input 212](img/vector212.png)

    input 212 -> angle 0.313725, length 0.823529

![input 213](img/vector213.png)

    input 213 -> angle 0.352941, length 0.823529

![input 214](img/vector214.png)

    input 214 -> angle 0.392157, length 0.823529

![input 215](img/vector215.png)

    input 215 -> angle 0.427451, length 0.823529

![input 216](img/vector216.png)

    input 216 -> angle 0.466667, length 0.823529

![input 217](img/vector217.png)

    input 217 -> angle 0.505882, length 0.823529

![input 218](img/vector218.png)

    input 218 -> angle 0.545098, length 0.823529

![input 219](img/vector219.png)

    input 219 -> angle 0.584314, length 0.823529

![input 220](img/vector220.png)

    input 220 -> angle 0.627451, length 0.862745

![input 221](img/vector221.png)

    input 221 -> angle 0.666667, length 0.862745

![input 222](img/vector222.png)

    input 222 -> angle 0.705882, length 0.862745

![input 223](img/vector223.png)

    input 223 -> angle 0.745098, length 0.862745

![input 224](img/vector224.png)

    input 224 -> angle 0.784314, length 0.862745

![input 225](img/vector225.png)

    input 225 -> angle 0.823529, length 0.862745

![input 226](img/vector226.png)

    input 226 -> angle 0.862745, length 0.862745

![input 227](img/vector227.png)

    input 227 -> angle 0.901961, length 0.862745

![input 228](img/vector228.png)

    input 228 -> angle 0.941176, length 0.862745

![input 229](img/vector229.png)

    input 229 -> angle 0.980392, length 0.862745

![input 230](img/vector230.png)

    input 230 -> angle 0.015686, length 0.901961

![input 231](img/vector231.png)

    input 231 -> angle 0.054902, length 0.901961

![input 232](img/vector232.png)

    input 232 -> angle 0.094118, length 0.901961

![input 233](img/vector233.png)

    input 233 -> angle 0.133333, length 0.901961

![input 234](img/vector234.png)

    input 234 -> angle 0.172549, length 0.901961

![input 235](img/vector235.png)

    input 235 -> angle 0.211765, length 0.901961

![input 236](img/vector236.png)

    input 236 -> angle 0.254902, length 0.901961

![input 237](img/vector237.png)

    input 237 -> angle 0.294118, length 0.901961

![input 238](img/vector238.png)

    input 238 -> angle 0.333333, length 0.901961

![input 239](img/vector239.png)

    input 239 -> angle 0.372549, length 0.901961

![input 240](img/vector240.png)

    input 240 -> angle 0.411765, length 0.941176

![input 241](img/vector241.png)

    input 241 -> angle 0.450980, length 0.941176

![input 242](img/vector242.png)

    input 242 -> angle 0.490196, length 0.941176

![input 243](img/vector243.png)

    input 243 -> angle 0.529412, length 0.941176

![input 244](img/vector244.png)

    input 244 -> angle 0.568627, length 0.941176

![input 245](img/vector245.png)

    input 245 -> angle 0.603922, length 0.941176

![input 246](img/vector246.png)

    input 246 -> angle 0.643137, length 0.941176

![input 247](img/vector247.png)

    input 247 -> angle 0.682353, length 0.941176

![input 248](img/vector248.png)

    input 248 -> angle 0.721569, length 0.941176

![input 249](img/vector249.png)

    input 249 -> angle 0.760784, length 0.941176

![input 250](img/vector250.png)

    input 250 -> angle 0.800000, length 0.980392

![input 251](img/vector251.png)

    input 251 -> angle 0.839216, length 0.980392

![input 252](img/vector252.png)

    input 252 -> angle 0.882353, length 0.980392

![input 253](img/vector253.png)

    input 253 -> angle 0.921569, length 0.980392

![input 254](img/vector254.png)

    input 254 -> angle 0.960784, length 0.980392

![input 255](img/vector255.png)

    input 255 -> angle 0.000000, length 0.980392
