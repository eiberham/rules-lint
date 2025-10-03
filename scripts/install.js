const fs = require('fs')
const path = require('path')
const os = require('os')
const { spawnSync } = require('child_process')

function getPlatformInfo() {
  const platform = os.platform()
  const arch = os.arch()
  
  const platforms = {
    'darwin': 'darwin',
    'linux': 'linux', 
    'win32': 'windows'
  }
  
  const archs = {
    'x64': 'amd64',
    'arm64': 'arm64'
  }
  
  return {
    platform: platforms[platform],
    arch: archs[arch]
  }
}

function install() {
  const { platform, arch } = getPlatformInfo()
  const binaryName = platform === 'windows' ? 'rules-lint.exe' : 'rules-lint'
  const sourcePath = path.join(__dirname, '..', 'dist', `rules-lint-${platform}-${arch}${platform === 'windows' ? '.exe' : ''}`)
  const binDir = path.join(__dirname, '..', 'bin')
  const binPath = path.join(binDir, binaryName)
  
  console.log(`🔍 Installing for ${platform}-${arch}...`)
  
  if (!fs.existsSync(binDir)) {
    fs.mkdirSync(binDir, { recursive: true })
  }
  
  if (!fs.existsSync(sourcePath)) {
    console.error(`❌ No pre-built binary found for ${platform}-${arch}`)
    console.error(`   Expected at: ${sourcePath}`)
    console.error(`   Trying to build from source...`)
    
    const buildResult = spawnSync('go', ['build', '-o', binPath, './cmd/lint'], {
      cwd: path.join(__dirname, '..'),
      stdio: 'inherit'
    })
    
    if (buildResult.status !== 0) {
      console.error(`❌ Build failed. Please ensure Go is installed.`)
      process.exit(1)
    }
    
    console.log(`✅ Built from source successfully`)
  } else {
    fs.copyFileSync(sourcePath, binPath)
    
    if (platform !== 'windows') {
      fs.chmodSync(binPath, '755')
    }
    
    console.log(`✅ Installed rules-lint for ${platform}-${arch}`)
  }
}

install()